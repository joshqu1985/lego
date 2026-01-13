package broker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/utils/routine"
)

const (
	RedisLogPrefix = "redis stream"
)

type (
	redisConsumer struct {
		wg              *sync.WaitGroup
		client          *redis.Client
		topics          map[string]string
		consumers       map[string]registered
		queue           chan *Message
		stopClaim       chan struct{}
		stopDelay       chan struct{}
		stopXread       chan struct{}
		stopWork        chan struct{}
		groupId         string
		name            string
		StreamMaxLength int64
		Concurrency     int
	}

	registered struct {
		callback ConsumeCallback
		offset   string
	}
)

var (
	DefaultBufferSize      = 100
	DefaultBlockTimeout    = 5 * time.Second
	DefaultClaimInterval   = 1 * time.Second
	DefaultDelayInterval   = 1 * time.Second
	DefaultMinIdle         = 30 * time.Minute
	DefaultWorkConcurrency = 5
	DefaultMaxRetryCount   = 5

	zrangeScript = redis.NewScript(`
local values = redis.call("zrangebyscore", KEYS[1], "-inf", ARGV[1], "LIMIT", "0", "10")

if not values or not values[1] then
    return nil
end

redis.call("zrem", KEYS[1], unpack(values))
return values
	`)

	hgetScript = redis.NewScript(`
local hashField = ARGV[1]
local value = redis.call("hget", KEYS[1], hashField)

if not value then
    return nil
end

redis.call("hdel", KEYS[1], hashField)
return value
	`)
)

func NewRedisConsumer(conf *Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Endpoints[0],
		Username: conf.AccessKey,
		Password: conf.SecretKey,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	hostname, _ := os.Hostname()

	return &redisConsumer{
		name:      hostname,
		client:    client,
		topics:    conf.Topics,
		consumers: make(map[string]registered),
		groupId:   conf.GroupId,
		queue:     make(chan *Message, DefaultBufferSize),
		wg:        &sync.WaitGroup{},

		stopClaim: make(chan struct{}, 1),
		stopDelay: make(chan struct{}, 1),
		stopXread: make(chan struct{}, 1),
		stopWork:  make(chan struct{}, DefaultWorkConcurrency),

		StreamMaxLength: MaxMessageCount,
		Concurrency:     DefaultWorkConcurrency,
	}, nil
}

func (rc *redisConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := rc.topics[topicKey]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}

	rc.consumers[topicVal] = registered{callback: f, offset: "$"}

	return nil
}

func (rc *redisConsumer) Start() error {
	if len(rc.consumers) == 0 {
		return errors.New("none consumer registered")
	}

	streams := make([]string, 0)
	for stream, consumer := range rc.consumers {
		err := rc.client.XGroupCreateMkStream(context.Background(), stream, rc.groupId, consumer.offset).Err()
		if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
			return err
		}
		streams = append(streams, stream)
	}

	totalConsumers := len(rc.consumers)
	for i := 0; i < totalConsumers; i++ {
		streams = append(streams, ">")
	}

	routine.Go(func() { rc.claim() })
	routine.Go(func() { rc.delay() })
	routine.Go(func() { rc.xread(streams) })

	stop := signalHandler()
	go func() {
		<-stop
		_ = rc.Close()
	}()

	rc.wg.Add(rc.Concurrency)
	for i := 0; i < rc.Concurrency; i++ {
		routine.Go(func() { rc.work() })
	}
	rc.wg.Wait()

	return nil
}

func (rc *redisConsumer) Close() error {
	rc.stopClaim <- struct{}{}
	rc.stopDelay <- struct{}{}
	rc.stopXread <- struct{}{}
	for i := 0; i < rc.Concurrency; i++ {
		rc.stopWork <- struct{}{}
	}

	return nil
}

func (rc *redisConsumer) claim() {
	ticker := time.NewTicker(DefaultClaimInterval)
	defer ticker.Stop()

	for {
		select {
		case <-rc.stopClaim:
			return
		case <-ticker.C:
			for stream := range rc.consumers {
				rc.claimStream(stream)
			}
		}
	}
}

func (rc *redisConsumer) delay() {
	ticker := time.NewTicker(DefaultDelayInterval)
	defer ticker.Stop()

	for {
		select {
		case <-rc.stopDelay:
			return
		case <-ticker.C:
			for stream := range rc.consumers {
				rc.repostDelay(stream)
			}
		}
	}
}

func (rc *redisConsumer) xread(streams []string) {
	for {
		select {
		case <-rc.stopXread:
			return
		default:
			args := &redis.XReadGroupArgs{
				Consumer: rc.name,
				Group:    rc.groupId,
				Streams:  streams,
				Count:    int64(DefaultBufferSize - len(rc.queue)),
				Block:    DefaultBlockTimeout,
			}
			vals, err := rc.client.XReadGroup(context.Background(), args).Result()
			if err != nil {
				logs.Errorf(RedisLogPrefix+" XReadGroup err:%v", err)

				continue
			}

			for _, val := range vals {
				rc.enqueue(val.Stream, val.Messages)
			}
		}
	}
}

func (rc *redisConsumer) enqueue(topic string, msgs []redis.XMessage) {
	for _, m := range msgs {
		msg := &Message{
			Topic:     topic,
			MessageId: m.ID,
		}

		if v, ok := m.Values["properties"]; ok {
			data, _ := v.(string)
			_ = json.Unmarshal([]byte(data), &msg.Properties)
		}

		if v, ok := m.Values["payload"]; ok {
			data, _ := v.(string)
			msg.Payload = []byte(data)
		}
		rc.queue <- msg
	}
}

func (rc *redisConsumer) work() {
	defer rc.wg.Done()

	for {
		select {
		case msg := <-rc.queue:
			register, ok := rc.consumers[msg.Topic]
			if !ok || register.callback == nil {
				if err := rc.client.XAck(context.Background(), msg.Topic, rc.groupId, msg.MessageId).Err(); err != nil {
					logs.Errorf(RedisLogPrefix+" XAck err:%v", err)
				}

				continue
			}

			_ = routine.Safe(func() {
				_ = register.callback(context.Background(), msg)
				if err := rc.client.XAck(context.Background(), msg.Topic, rc.groupId, msg.MessageId).Err(); err != nil {
					logs.Errorf(RedisLogPrefix+" XAck err:%v", err)
				}
			})
		case <-rc.stopWork:
			return
		}
	}
}

func (rc *redisConsumer) claimStream(stream string) {
	start, end := "-", "+"
	for {
		pendings, err := rc.client.XPendingExt(context.Background(), &redis.XPendingExtArgs{
			Stream: stream,
			Group:  rc.groupId,
			Start:  start,
			End:    end,
			Count:  int64(DefaultBufferSize - len(rc.queue)),
		}).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logs.Errorf(RedisLogPrefix+" XPendingExt err:%v", err)

			break
		}
		if len(pendings) == 0 {
			break
		}

		for _, pending := range pendings {
			if pending.Idle < DefaultMinIdle {
				continue
			}

			claimMsgs, xerr := rc.client.XClaim(context.Background(), &redis.XClaimArgs{
				Consumer: rc.name,
				Stream:   stream,
				Group:    rc.groupId,
				MinIdle:  DefaultMinIdle,
				Messages: []string{pending.ID},
			}).Result()
			if xerr != nil && !errors.Is(xerr, redis.Nil) {
				logs.Errorf(RedisLogPrefix+" XClaim err:%v", err)

				break
			}
			if errors.Is(xerr, redis.Nil) {
				if yerr := rc.client.XAck(context.Background(), stream, rc.groupId, pending.ID).Err(); yerr != nil {
					logs.Errorf(RedisLogPrefix+" XAck err:%v", yerr)

					continue
				}
			}

			if pending.RetryCount > int64(DefaultMaxRetryCount) {
				for _, msg := range claimMsgs {
					if zerr := rc.client.XAck(context.Background(), stream, rc.groupId, msg.ID).Err(); zerr != nil {
						logs.Errorf(RedisLogPrefix+" XAck retry count gt N err:%v", zerr)

						continue
					}
				}

				continue
			}
			rc.enqueue(stream, claimMsgs)
		}

		next, err := nextMessageID(pendings[len(pendings)-1].ID)
		if err != nil {
			logs.Errorf(RedisLogPrefix+" calc next msg id err:%v", err)

			break
		}
		start = next
	}
}

func (rc *redisConsumer) repostDelay(stream string) {
	for {
		vals, err := zrangeScript.Run(context.Background(), rc.client,
			[]string{fmt.Sprintf(zsetFormat, stream)}, strconv.FormatInt(time.Now().Unix(), 10)).StringSlice()
		if err != nil && !errors.Is(err, redis.Nil) {
			logs.Errorf(RedisLogPrefix+" zrange delay messages err:%v", err)

			break
		}
		if len(vals) == 0 {
			break
		}

		hashkey := []string{fmt.Sprintf(hashFormat, stream)}
		for _, val := range vals {
			bytes, xerr := hgetScript.Run(context.Background(), rc.client, hashkey, val).Text()
			if xerr != nil {
				continue
			}

			msg := &Message{}
			if yerr := json.Unmarshal([]byte(bytes), msg); yerr != nil {
				continue
			}

			values := map[string]any{"payload": msg.Payload}
			if len(msg.Properties) != 0 {
				data, _ := json.Marshal(&msg.Properties)
				values["properties"] = data
			}
			if _, zerr := rc.client.XAdd(context.Background(), &redis.XAddArgs{
				Stream: stream,
				MaxLen: rc.StreamMaxLength,
				Values: values,
			}).Result(); zerr != nil {
				logs.Errorf(RedisLogPrefix+" XAdd delay msg err:%v", zerr)
			}
		}
	}
}

func nextMessageID(id string) (string, error) {
	parts := strings.Split(id, "-")
	index := parts[1]
	parsed, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%d", parts[0], parsed+1), nil
}

func signalHandler() <-chan struct{} {
	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		close(stop)
	}()

	return stop
}
