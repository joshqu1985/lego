package broker

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// NewKafkaProducer 创建Producer
func NewKafkaProducer(conf Config) (Producer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}

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

func (this *kafkaProducer) Send(ctx context.Context, topicKey string, msg *Message, delay ...int64) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}

	topicVal, ok := this.topics[topicKey]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	data := kafka.Message{
		Topic:   topicVal,
		Headers: []kafka.Header{},
		Key:     []byte(msg.Key),
		Value:   msg.Payload,
	}
	for key, val := range msg.Properties {
		data.Headers = append(data.Headers, kafka.Header{Key: key, Value: []byte(val)})
	}
	return this.client.WriteMessages(ctx, data)
}

func (this *kafkaProducer) Close() error {
	return this.client.Close()
}
