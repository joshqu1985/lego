package broker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/xid"
)

const (
	zsetFormat = "{%s_delay}.zset"
	hashFormat = "{%s_delay}.hash"
)

type redisProducer struct {
	client          *redis.Client
	topics          map[string]string
	StreamMaxLength int64
}

var delayScript = redis.NewScript(`
local key = ARGV[1]
local val = ARGV[2]
local score = tonumber(ARGV[3])

redis.call("zadd", KEYS[1], score, key)
redis.call("hset", KEYS[2], key, val)
return 1
	`)

func NewRedisProducer(conf *Config) (Producer, error) {
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

	return &redisProducer{
		client:          client,
		topics:          conf.Topics,
		StreamMaxLength: MaxMessageCount,
	}, nil
}

func (rp *redisProducer) Send(ctx context.Context, topic string, msg *Message) error {
	return rp.SendDelay(ctx, topic, msg, 0)
}

func (rp *redisProducer) SendDelay(ctx context.Context, topic string, msg *Message, stamp int64) error {
	if msg == nil {
		return errors.New(ErrMessageIsNil)
	}

	topicVal, ok := rp.topics[topic]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}
	msg.Topic = topicVal

	if rp.CountStream(ctx, topicVal)+rp.CountDelay(ctx, topicVal) > MaxMessageCount {
		return errors.New(ErrQueueIsFull)
	}

	if stamp > time.Now().Unix() {
		return rp.sendDelay(ctx, msg, stamp)
	}

	return rp.sendStream(ctx, msg)
}

func (rp *redisProducer) Close() error {
	return rp.client.Close()
}

func (rp *redisProducer) sendStream(ctx context.Context, msg *Message) error {
	values := map[string]any{"payload": msg.Payload}
	if len(msg.Properties) != 0 {
		data, _ := json.Marshal(&msg.Properties)
		values["properties"] = data
	}

	var err error
	args := &redis.XAddArgs{
		Stream: msg.Topic,
		MaxLen: rp.StreamMaxLength,
		Values: values,
	}
	msg.MessageId, err = rp.client.XAdd(ctx, args).Result()

	return err
}

func (rp *redisProducer) CountStream(ctx context.Context, topic string) int64 {
	return rp.client.XLen(ctx, topic).Val()
}

func (rp *redisProducer) sendDelay(ctx context.Context, msg *Message, timestamp int64) error {
	keys := []string{fmt.Sprintf(zsetFormat, msg.Topic), fmt.Sprintf(hashFormat, msg.Topic)}
	data, _ := json.Marshal(msg)
	vals := []any{xid.New().String(), data, timestamp}

	ok, err := delayScript.Run(ctx, rp.client, keys, vals...).Bool()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("duplicate message")
	}

	return nil
}

func (rp *redisProducer) CountDelay(ctx context.Context, topic string) int64 {
	return rp.client.ZCard(ctx, fmt.Sprintf(zsetFormat, topic)).Val()
}
