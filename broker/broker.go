package broker

import (
	"context"
	"fmt"
)

// Producer 生产者
type Producer interface {
	Send(ctx context.Context, topic string, msg *Message, delay ...int64) error // delay 时间戳
	Close() error
}

// Consumer 消费者
type Consumer interface {
	Register(topic string, f ConsumeCallback) error
	Start() error
	Close() error
}

// NewProducer 创建Producer
func NewProducer(conf Config) Producer {
	var (
		producer Producer
		err      error
	)

	switch conf.Source {
	case "kafka":
		producer, err = NewKafkaProducer(conf)
	case "rocketmq":
		producer, err = NewRocketProducer(conf)
	case "pulsar":
		producer, err = NewPulsarProducer(conf)
	case "redis":
		producer, err = NewRedisProducer(conf)
	case "memory":
		producer, err = NewMemoryProducer(conf)
	default:
		producer, err = nil, fmt.Errorf("unknown source: %s", conf.Source)
	}

	if err != nil {
		panic(err)
	}
	return producer
}

// ConsumeCallback 消费回调
type ConsumeCallback func(context.Context, *Message) error

// NewConsumer 创建Consumer
func NewConsumer(conf Config) Consumer {
	var (
		consumer Consumer
		err      error
	)

	switch conf.Source {
	case "kafka":
		consumer, err = NewKafkaConsumer(conf)
	case "rocketmq":
		consumer, err = NewRocketConsumer(conf)
	case "pulsar":
		consumer, err = NewPulsarConsumer(conf)
	case "redis":
		consumer, err = NewRedisConsumer(conf)
	case "memory":
		consumer, err = NewMemoryConsumer(conf)
	default:
		consumer, err = nil, fmt.Errorf("unknown source: %s", conf.Source)
	}

	if err != nil {
		panic(err)
	}
	return consumer
}
