package broker

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"

	"github.com/joshqu1985/lego/utils/routine"
)

type kafkaConsumer struct {
	topics    map[string]string
	callbacks map[string]ConsumeCallback
	client    *kafka.Reader
	stopWork  chan struct{}
	groupId   string
}

// NewKafkaConsumer 创建Consumer.
func NewKafkaConsumer(conf *Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
	}

	topicVals := make([]string, 0)
	for _, topicVal := range conf.Topics {
		topicVals = append(topicVals, topicVal)
	}

	client := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     conf.Endpoints,
		GroupID:     conf.GroupId,
		GroupTopics: topicVals,
		StartOffset: kafka.LastOffset,
	})

	return &kafkaConsumer{
		topics:    conf.Topics,
		groupId:   conf.GroupId,
		client:    client,
		callbacks: make(map[string]ConsumeCallback),
		stopWork:  make(chan struct{}, 1),
	}, nil
}

func (kc *kafkaConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := kc.topics[topicKey]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}

	kc.callbacks[topicVal] = f

	return nil
}

func (kc *kafkaConsumer) Start() error {
	for {
		select {
		case <-kc.stopWork:
			return nil
		default:
			msg, err := kc.client.ReadMessage(context.Background())
			if err != nil {
				continue
			}

			data := &Message{
				Topic:      msg.Topic,
				Properties: make(map[string]string),
				Payload:    msg.Value,
			}
			for _, header := range msg.Headers {
				data.Properties[header.Key] = string(header.Value)
			}
			fn, ok := kc.callbacks[msg.Topic]
			if !ok || fn == nil {
				continue
			}

			_ = routine.Safe(func() {
				_ = fn(context.Background(), data)
			})
		}
	}
}

func (kc *kafkaConsumer) Close() error {
	kc.stopWork <- struct{}{}

	return kc.client.Close()
}
