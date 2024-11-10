package broker

import (
	"context"
	"errors"
)

// Message 消息结构
type Message struct {
	Properties map[string]string
	Payload    []byte
	MessageId  string // 不需要传 发送后会自动填入
	Topic      string // 不需要传 内部使用
}

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

var (
	MaxMessageCount = int64(500)
)

const (
	KAFKA  = 1 // kafka
	ROCKET = 2 // rocketmq
	PULSAR = 3 // pulsar
	REDIS  = 4 // redis stream
	MEMORY = 5 // memory
)

// ErrUnknowType 暂时不支持的类型错误
var ErrUnknowType = errors.New("unknow broker type")

// ConsumeCallback 消费回调
type ConsumeCallback func(context.Context, *Message) error

// NewProducer 创建Producer
func NewProducer(broker int, conf Config) Producer {
	var (
		producer Producer
		err      error
	)

	switch broker {
	case KAFKA:
		producer, err = NewKafkaProducer(conf)
	case ROCKET:
		producer, err = NewRocketProducer(conf)
	case REDIS:
		producer, err = NewRedisProducer(conf)
	case MEMORY:
		producer, err = NewMemoryProducer(conf)
	default:
		producer, err = nil, ErrUnknowType
	}

	if err != nil {
		panic(err)
	}
	return producer
}

// NewConsumer 创建Consumer
func NewConsumer(broker int, conf Config) Consumer {
	var (
		consumer Consumer
		err      error
	)

	switch broker {
	case KAFKA:
		consumer, err = NewKafkaConsumer(conf)
	case ROCKET:
		consumer, err = NewRocketConsumer(conf)
	case REDIS:
		consumer, err = NewRedisConsumer(conf)
	case MEMORY:
		consumer, err = NewMemoryConsumer(conf)
	default:
		consumer, err = nil, ErrUnknowType
	}

	if err != nil {
		panic(err)
	}
	return consumer
}
