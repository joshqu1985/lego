package broker

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func NewPulsarProducer(conf Config) (Producer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: conf.Endpoints[0],
	})
	if err != nil {
		return nil, err
	}

	producers := map[string]pulsar.Producer{}
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

type pulsarProducer struct {
	client    pulsar.Client
	producers map[string]pulsar.Producer
	topics    map[string]string
}

func (this *pulsarProducer) Send(ctx context.Context, topicKey string, msg *Message, args ...int64) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}

	producer, ok := this.producers[topicKey]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	data := pulsar.ProducerMessage{
		Payload:    msg.Payload,
		Key:        msg.Key,
		Properties: msg.Properties,
	}

	if len(args) != 0 && args[0] > time.Now().Unix() {
		data.DeliverAt = time.Unix(args[0], 0)
	}

	recive, err := producer.Send(ctx, &data)
	if recive != nil {
		msg.MessageId = recive.String()
	}
	return err
}

func (this *pulsarProducer) Close() error {
	for _, producer := range this.producers {
		producer.Close()
	}
	this.client.Close()
	return nil
}
