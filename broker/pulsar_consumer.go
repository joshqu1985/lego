package broker

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/pulsar-client-go/pulsar"
)

// NewPulsarConsumer 创建PulsarConsumer
func NewPulsarConsumer(conf Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: conf.Endpoints[0],
	})
	if err != nil {
		return nil, err
	}

	consumer := &pulsarConsumer{
		client:    client,
		topics:    conf.Topics,
		channel:   make(chan pulsar.ConsumerMessage, 100),
		callbacks: make(map[string]ConsumeCallback),
	}

	topicVals := []string{}
	for _, topicVal := range conf.Topics {
		topicVals = append(topicVals, topicVal)
	}

	options := pulsar.ConsumerOptions{
		SubscriptionName: conf.GroupId,
		Topics:           topicVals,
		Type:             pulsar.Shared,
		MessageChannel:   consumer.channel,
	}
	consumer.consumer, err = client.Subscribe(options)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// pulsarConsumer pulsar消费者结构
type pulsarConsumer struct {
	client    pulsar.Client
	consumer  pulsar.Consumer
	topics    map[string]string
	channel   chan pulsar.ConsumerMessage
	callbacks map[string]ConsumeCallback
}

func (this *pulsarConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := this.topics[topicKey]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	this.callbacks[topicVal] = f
	return nil
}

func (this *pulsarConsumer) Start() error {
	if len(this.callbacks) == 0 {
		return fmt.Errorf("at least one consumer function registered")
	}

	for cm := range this.channel {
		msg := cm.Message
		vals := strings.Split(msg.Topic(), "-partition-")
		if len(vals) != 2 {
			continue
		}
		callback, ok := this.callbacks[vals[0]]
		if !ok {
			continue
		}

		data := &Message{
			Topic:      vals[0],
			Properties: msg.Properties(),
			Payload:    msg.Payload(),
		}
		if err := callback(context.Background(), data); err == nil {
			cm.Consumer.Ack(msg)
		} else {
			cm.Consumer.Nack(msg)
		}
	}
	return nil
}

func (this *pulsarConsumer) Close() error {
	this.consumer.Close()
	this.client.Close()
	return nil
}
