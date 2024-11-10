package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/xid"
)

func NewRedisProducer(conf Config) (Producer, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Endpoints[0],
		Username: conf.AccessKey,
		Password: conf.SecretKey,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &redisProducer{
		client:          client,
		topics:          conf.Topics,
		StreamMaxLength: MaxMessageCount,
	}, nil
}

type redisProducer struct {
	client          *redis.Client
	topics          map[string]string
	StreamMaxLength int64
}

func (this *redisProducer) Send(ctx context.Context, topic string, msg *Message, args ...int64) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}
	if v, ok := this.topics[topic]; !ok {
		return fmt.Errorf("topic not found")
	} else {
		msg.Topic = v
	}

	if len(args) != 0 && args[0] > time.Now().Unix() {
		return this.sendDelay(ctx, msg, args[0])
	}

	return this.sendStream(ctx, msg)
}

func (this *redisProducer) Close() error {
	return nil
}

func (this *redisProducer) sendStream(ctx context.Context, msg *Message) error {
	values := map[string]any{"payload": msg.Payload}
	if len(msg.Properties) != 0 {
		data, _ := json.Marshal(&msg.Properties)
		values["properties"] = data
	}

	var err error
	msg.MessageId, err = this.client.XAdd(ctx, &redis.XAddArgs{
		Stream: msg.Topic,
		MaxLen: this.StreamMaxLength,
		Values: values,
	}).Result()
	return err
}

var (
	delayScript = redis.NewScript(`
local key = ARGV[1]
local val = ARGV[2]
local score = tonumber(ARGV[3])

redis.call("zadd", KEYS[1], score, key)
redis.call("hset", KEYS[2], key, val)
return 1
	`)
)

const (
	zsetFormat = "{%s_delay}.zset"
	hashFormat = "{%s_delay}.hash"
)

func (this *redisProducer) sendDelay(ctx context.Context, msg *Message, timestamp int64) error {
	data, _ := json.Marshal(msg)
	ok, err := delayScript.Run(ctx, this.client,
		[]string{fmt.Sprintf(zsetFormat, msg.Topic), fmt.Sprintf(hashFormat, msg.Topic)},
		xid.New().String(), data, timestamp).Bool()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("duplicate message")
	}
	return nil
}
