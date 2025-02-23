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

	topicVal, ok := this.topics[topic]
	if !ok {
		return fmt.Errorf("topic not found")
	}
	msg.Topic = topicVal

	if this.CountStream(ctx, topicVal)+this.CountDelay(ctx, topicVal) > MaxMessageCount {
		return fmt.Errorf("queue is full")
	}

	if len(args) != 0 && args[0] > time.Now().Unix() {
		return this.SendDelay(ctx, msg, args[0])
	}
	return this.SendStream(ctx, msg)
}

func (this *redisProducer) Close() error {
	return nil
}

func (this *redisProducer) SendStream(ctx context.Context, msg *Message) error {
	values := map[string]any{"payload": msg.Payload}
	if len(msg.Properties) != 0 {
		data, _ := json.Marshal(&msg.Properties)
		values["properties"] = data
	}

	var err error
	args := &redis.XAddArgs{
		Stream: msg.Topic,
		MaxLen: this.StreamMaxLength,
		Values: values,
	}
	msg.MessageId, err = this.client.XAdd(ctx, args).Result()
	return err
}

func (this *redisProducer) CountStream(ctx context.Context, topic string) int64 {
	return this.client.XLen(ctx, topic).Val()
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

func (this *redisProducer) SendDelay(ctx context.Context, msg *Message, timestamp int64) error {
	keys := []string{fmt.Sprintf(zsetFormat, msg.Topic), fmt.Sprintf(hashFormat, msg.Topic)}
	data, _ := json.Marshal(msg)
	vals := []any{xid.New().String(), data, timestamp}

	ok, err := delayScript.Run(ctx, this.client, keys, vals...).Bool()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("duplicate message")
	}
	return nil
}

func (this *redisProducer) CountDelay(ctx context.Context, topic string) int64 {
	return this.client.ZCard(ctx, fmt.Sprintf(zsetFormat, topic)).Val()
}
