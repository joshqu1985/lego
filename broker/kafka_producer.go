package broker

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"
)

type kafkaProducer struct {
	client *kafka.Writer
	topics map[string]string
}

// NewKafkaProducer 创建Producer.
func NewKafkaProducer(conf *Config) (Producer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
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

func (kp *kafkaProducer) Send(ctx context.Context, topicKey string, msg *Message) error {
	if msg == nil {
		return errors.New(ErrMessageIsNil)
	}

	topicVal, ok := kp.topics[topicKey]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}

	data := kafka.Message{
		Topic:   topicVal,
		Headers: make([]kafka.Header, 0),
		Key:     []byte(msg.Key),
		Value:   msg.Payload,
	}
	for key, val := range msg.Properties {
		data.Headers = append(data.Headers, kafka.Header{Key: key, Value: []byte(val)})
	}

	return kp.client.WriteMessages(ctx, data)
}

func (kp *kafkaProducer) SendDelay(ctx context.Context, topicKey string, msg *Message, stamp int64) error {
	return errors.New("not support")
}

func (kp *kafkaProducer) Close() error {
	return kp.client.Close()
}
