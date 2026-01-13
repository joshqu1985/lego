package broker

import (
	"context"
	"errors"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type pulsarProducer struct {
	client    pulsar.Client
	producers map[string]pulsar.Producer
	topics    map[string]string
}

func NewPulsarProducer(conf *Config) (Producer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: conf.Endpoints[0],
	})
	if err != nil {
		return nil, err
	}

	producers := make(map[string]pulsar.Producer)
	for topicKey, topicVal := range conf.Topics {
		options := pulsar.ProducerOptions{
			Topic: topicVal,
		}
		producers[topicKey], err = client.CreateProducer(options)
		if err != nil {
			return nil, err
		}
	}

	return &pulsarProducer{
		client:    client,
		producers: producers,
		topics:    conf.Topics,
	}, nil
}

func (pp *pulsarProducer) Send(ctx context.Context, topicKey string, msg *Message) error {
	return pp.SendDelay(ctx, topicKey, msg, 0)
}

func (pp *pulsarProducer) SendDelay(ctx context.Context, topicKey string, msg *Message, stamp int64) error {
	if msg == nil {
		return errors.New(ErrMessageIsNil)
	}

	producer, ok := pp.producers[topicKey]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}

	data := pulsar.ProducerMessage{
		Payload:    msg.Payload,
		Key:        msg.Key,
		Properties: msg.Properties,
	}

	if stamp > time.Now().Unix() {
		data.DeliverAt = time.Unix(stamp, 0)
	}

	recive, err := producer.Send(ctx, &data)
	if recive != nil {
		msg.MessageId = recive.String()
	}

	return err
}

func (pp *pulsarProducer) Close() error {
	for _, producer := range pp.producers {
		producer.Close()
	}
	pp.client.Close()

	return nil
}
