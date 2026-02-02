package broker

import (
	"context"
	"errors"
)

const (
	SourceKafka   = "kafka"
	SourceRocket  = "rocketmq"
	SourcePulsar  = "pulsar"
	SourceRedis   = "redis"
	SourceMemory  = "memory"
	SourceUnknown = "unsupported"
)

type (
	// Producer 生产者.
	Producer interface {
		Send(ctx context.Context, topic string, msg *Message) error
		SendDelay(ctx context.Context, topic string, msg *Message, stamp int64) error // stamp 时间戳
		Close() error
	}

	// Consumer 消费者.
	Consumer interface {
		Register(topic string, f ConsumeCallback) error
		Start() error
		Close() error
	}

	// ConsumeCallback 消费回调.
	ConsumeCallback func(context.Context, *Message) error
)

var (
	ErrEndpointsEmpty   = errors.New("endpoints is empty")
	ErrSubscriberNil    = errors.New("subscriber is nil")
	ErrMessageIsNil     = errors.New("message is nil")
	ErrQueueIsFull      = errors.New("queue is full")
	ErrTopicNotFound    = errors.New("topic not found")
	ErrSourceNotSupport = errors.New("source not support")
)

// NewProducer 创建Producer.
func NewProducer(conf *Config) (Producer, error) {
	switch conf.Source {
	case SourceKafka:
		return NewKafkaProducer(conf)
	case SourceRocket:
		return NewRocketProducer(conf)
	case SourcePulsar:
		return NewPulsarProducer(conf)
	case SourceRedis:
		return NewRedisProducer(conf)
	case SourceMemory:
		return NewMemoryProducer(conf)
	default:
		return nil, ErrSourceNotSupport
	}
}

// NewConsumer 创建Consumer.
func NewConsumer(conf *Config) (Consumer, error) {
	switch conf.Source {
	case SourceKafka:
		return NewKafkaConsumer(conf)
	case SourceRocket:
		return NewRocketConsumer(conf)
	case SourcePulsar:
		return NewPulsarConsumer(conf)
	case SourceRedis:
		return NewRedisConsumer(conf)
	case SourceMemory:
		return NewMemoryConsumer(conf)
	default:
		return nil, ErrSourceNotSupport
	}
}
