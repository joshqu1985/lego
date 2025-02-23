package broker

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// NewKafkaConsumer 创建Consumer
func NewKafkaConsumer(conf Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}

	topicVals := []string{}
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
		ctx:       context.Background(),
		client:    client,
		callbacks: map[string]ConsumeCallback{},
	}, nil
}

type kafkaConsumer struct {
	topics    map[string]string
	groupId   string
	callbacks map[string]ConsumeCallback
	ctx       context.Context
	client    *kafka.Reader
}

func (this *kafkaConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := this.topics[topicKey]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	this.callbacks[topicVal] = f
	return nil
}

func (this *kafkaConsumer) Start() error {
	for {
		msg, err := this.client.ReadMessage(this.ctx)
		if err != nil {
			continue
		}

		data := &Message{
			Topic:      msg.Topic,
			Properties: map[string]string{},
			Payload:    msg.Value,
		}
		for _, header := range msg.Headers {
			data.Properties[header.Key] = string(header.Value)
		}
		callback, ok := this.callbacks[msg.Topic]
		if !ok || callback == nil {
			continue
		}

		callback(context.Background(), data)
	}
}

func (this *kafkaConsumer) Close() error {
	this.ctx.Done()
	return this.client.Close()
}
