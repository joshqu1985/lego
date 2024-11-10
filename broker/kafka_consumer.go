package broker

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// NewKafkaConsumer 创建Consumer
func NewKafkaConsumer(conf Config) (*kafkaConsumer, error) {
	topics := []string{}
	for _, topic := range conf.Topics {
		topics = append(topics, topic)
	}
	client := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     conf.Endpoints,
		GroupID:     conf.GroupId,
		GroupTopics: topics,
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

func (this *kafkaConsumer) Register(topic string, f ConsumeCallback) error {
	realTopic := ""
	if v, ok := this.topics[topic]; !ok {
		return fmt.Errorf("topic not found")
	} else {
		realTopic = v
	}

	this.callbacks[realTopic] = f
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
