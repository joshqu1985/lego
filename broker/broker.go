package broker

import (
	"context"
	"errors"
)

// Message 消息结构
type Message struct {
	Properties map[string]string
	Key        string
	Payload    []byte
	MessageId  string // 不需要传 发送后会自动填入
	Topic      string // 不需要传 内部使用
}

func (this *Message) SetTag(tag string) {
	if this.Properties == nil {
		this.Properties = make(map[string]string)
	}
	this.Properties["tag"] = tag
}

func (this *Message) GetTag() string {
	tag, ok := this.Properties["tag"]
	if !ok {
		return ""
	}
	return tag
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
	MaxMessageCount = int64(10000)
)

// ErrUnknowType 暂时不支持的类型错误
var ErrUnknowType = errors.New("unknow broker type")

// ConsumeCallback 消费回调
type ConsumeCallback func(context.Context, *Message) error

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
		producer, err = nil, ErrUnknowType
	}

	if err != nil {
		panic(err)
	}
	return producer
}

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
		consumer, err = nil, ErrUnknowType
	}

	if err != nil {
		panic(err)
	}
	return consumer
}
