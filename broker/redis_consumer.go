package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/redis/go-redis/v9"

	"github.com/joshqu1985/lego/utils/routine"
)

var (
	DefaultBufferSize      = 100
	DefaultBlockTimeout    = 5 * time.Second
	DefaultClaimInterval   = 1 * time.Second
	DefaultDelayInterval   = 1 * time.Second
	DefaultMinIdle         = 30 * time.Minute
	DefaultWorkConcurrency = 5
	DefaultMaxRetryCount   = 5
)

func NewRedisConsumer(conf Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
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

type redisConsumer struct {
	name      string
	client    *redis.Client
	topics    map[string]string
	consumers map[string]registered
	groupId   string
	queue     chan *Message
	wg        *sync.WaitGroup

	stopClaim chan struct{}
	stopDelay chan struct{}
	stopXread chan struct{}
	stopWork  chan struct{}

	StreamMaxLength int64
	Concurrency     int
}

func (this *redisConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := this.topics[topicKey]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	this.consumers[topicVal] = registered{callback: f, offset: "$"}
	return nil
}

func (this *redisConsumer) Start() error {
	if len(this.consumers) == 0 {
		return fmt.Errorf("none consumer registered")
	}

	streams := []string{}
	for stream, consumer := range this.consumers {
		err := this.client.XGroupCreateMkStream(context.Background(), stream, this.groupId, consumer.offset).Err()
		if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
			return err
		}
		streams = append(streams, stream)
	}
	for i := 0; i < len(this.consumers); i++ {
		streams = append(streams, ">")
	}

	routine.Go(func() { this.claim() })
	routine.Go(func() { this.delay() })
	routine.Go(func() { this.xread(streams) })

	stop := signalHandler()
	go func() {
		<-stop
		_ = this.Close()
	}()

	this.wg.Add(this.Concurrency)
	for i := 0; i < this.Concurrency; i++ {
		routine.Go(func() { this.work() })
	}
	this.wg.Wait()

	return nil
}

func (this *redisConsumer) Close() error {
	this.stopClaim <- struct{}{}
	this.stopDelay <- struct{}{}
	this.stopXread <- struct{}{}
	for i := 0; i < this.Concurrency; i++ {
		this.stopWork <- struct{}{}
	}
	return nil
}

type registered struct {
	callback ConsumeCallback
	offset   string
}

func (this *redisConsumer) claim() {
	ticker := time.NewTicker(DefaultClaimInterval)
	defer ticker.Stop()

	for {
		select {
		case <-this.stopClaim:
			return
		case <-ticker.C:
			for stream := range this.consumers {
				this.claimStream(stream)
			}
		}
	}
}

func (this *redisConsumer) delay() {
	ticker := time.NewTicker(DefaultDelayInterval)
	defer ticker.Stop()

	for {
		select {
		case <-this.stopDelay:
			return
		case <-ticker.C:
			for stream := range this.consumers {
				this.repostDelay(stream)
			}
		}
	}
}

func (this *redisConsumer) xread(streams []string) {
	for {
		select {
		case <-this.stopXread:
			return
		default:
			args := &redis.XReadGroupArgs{
				Consumer: this.name,
				Group:    this.groupId,
				Streams:  streams,
				Count:    int64(DefaultBufferSize - len(this.queue)),
				Block:    DefaultBlockTimeout,
			}
			vals, err := this.client.XReadGroup(context.Background(), args).Result()
			if err != nil {
				glog.Errorf("redis stream XReadGroup err:%v", err)
				continue
			}

			for _, val := range vals {
				this.enqueue(val.Stream, val.Messages)
			}
		}
	}
}

func (this *redisConsumer) enqueue(topic string, msgs []redis.XMessage) {
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
		this.queue <- msg
	}
}

func (this *redisConsumer) work() {
	defer this.wg.Done()

	for {
		select {
		case msg := <-this.queue:
			register, ok := this.consumers[msg.Topic]
			if !ok || register.callback == nil {
				if err := this.client.XAck(context.Background(), msg.Topic, this.groupId, msg.MessageId).Err(); err != nil {
					glog.Errorf("redis stream claim err:%v", err)
				}
				continue
			}

			routine.Safe(func() {
				_ = register.callback(context.Background(), msg)
				if err := this.client.XAck(context.Background(), msg.Topic, this.groupId, msg.MessageId).Err(); err != nil {
					glog.Errorf("redis stream claim err:%v", err)
				}
			})
		case <-this.stopWork:
			return
		}
	}
}

func (this *redisConsumer) claimStream(stream string) {
	start, end := "-", "+"
	for {
		pendings, err := this.client.XPendingExt(context.Background(), &redis.XPendingExtArgs{
			Stream: stream,
			Group:  this.groupId,
			Start:  start,
			End:    end,
			Count:  int64(DefaultBufferSize - len(this.queue)),
		}).Result()
		if err != nil && err != redis.Nil {
			glog.Errorf("redis stream list pending messages err:%v", err)
			break
		}
		if len(pendings) == 0 {
			break
		}

		for _, pending := range pendings {
			if pending.Idle < DefaultMinIdle {
				continue
			}

			claimMsgs, err := this.client.XClaim(context.Background(), &redis.XClaimArgs{
				Consumer: this.name,
				Stream:   stream,
				Group:    this.groupId,
				MinIdle:  DefaultMinIdle,
				Messages: []string{pending.ID},
			}).Result()
			if err != nil && err != redis.Nil {
				glog.Errorf("redis stream claim err:%v", err)
				break
			}
			if err == redis.Nil {
				if err := this.client.XAck(context.Background(), stream, this.groupId, pending.ID).Err(); err != nil {
					glog.Errorf("redis stream claim ack err:%v", err)
					continue
				}
			}

			if pending.RetryCount > int64(DefaultMaxRetryCount) {
				for _, msg := range claimMsgs {
					if err := this.client.XAck(context.Background(), stream, this.groupId, msg.ID).Err(); err != nil {
						glog.Errorf("redis stream ack retry count gt N err:%v", err)
						continue
					}
				}
				continue
			}
			this.enqueue(stream, claimMsgs)
		}

		next, err := nextMessageID(pendings[len(pendings)-1].ID)
		if err != nil {
			glog.Errorf("redis stream calc next msg id err:%v", err)
			break
		}
		start = next
	}
}

var (
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

func (this *redisConsumer) repostDelay(stream string) {
	for {
		vals, err := zrangeScript.Run(context.Background(), this.client,
			[]string{fmt.Sprintf(zsetFormat, stream)}, fmt.Sprintf("%d", time.Now().Unix())).StringSlice()
		if err != nil && err != redis.Nil {
			glog.Errorf("redis stream zrange delay messages err:%v", err)
			break
		}
		if len(vals) == 0 {
			break
		}

		hashkey := []string{fmt.Sprintf(hashFormat, stream)}
		for _, val := range vals {
			bytes, err := hgetScript.Run(context.Background(), this.client, hashkey, val).Text()
			if err != nil {
				continue
			}

			msg := &Message{}
			if err := json.Unmarshal([]byte(bytes), msg); err != nil {
				continue
			}

			values := map[string]any{"payload": msg.Payload}
			if len(msg.Properties) != 0 {
				data, _ := json.Marshal(&msg.Properties)
				values["properties"] = data
			}
			if _, err := this.client.XAdd(context.Background(), &redis.XAddArgs{
				Stream: stream,
				MaxLen: this.StreamMaxLength,
				Values: values,
			}).Result(); err != nil {
				glog.Errorf("redis stream xadd delay msg err:%v", err)
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
		<-c
		os.Exit(1)
	}()
	return stop
}
