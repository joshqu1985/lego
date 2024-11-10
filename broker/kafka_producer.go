package broker

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// NewKafkaProducer 创建Producer
func NewKafkaProducer(conf Config) (*kafkaProducer, error) {
	client := kafka.NewWriter(kafka.WriterConfig{
		Brokers: conf.Endpoints,
		Async:   true,
	})
	return &kafkaProducer{
		client: client,
		topics: conf.Topics,
	}, nil
}

type kafkaProducer struct {
	client *kafka.Writer
	topics map[string]string
}

func (this *kafkaProducer) Send(ctx context.Context, topic string, msg *Message, delay ...int64) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}
	realTopic, ok := this.topics[topic]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	data := kafka.Message{
		Topic:   realTopic,
		Headers: []kafka.Header{},
		Value:   msg.Payload,
	}
	for key, val := range msg.Properties {
		data.Headers = append(data.Headers, kafka.Header{
			Key:   key,
			Value: []byte(val),
		})
	}
	return this.client.WriteMessages(ctx, data)
}

func (this *kafkaProducer) Close() error {
	return this.client.Close()
}
