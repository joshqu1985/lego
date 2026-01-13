package broker

import (
	"context"
	"errors"
	"strings"

	"github.com/apache/pulsar-client-go/pulsar"

	"github.com/joshqu1985/lego/utils/routine"
)

// pulsarConsumer pulsar消费者结构.
type pulsarConsumer struct {
	client    pulsar.Client
	consumer  pulsar.Consumer
	topics    map[string]string
	channel   chan pulsar.ConsumerMessage
	callbacks map[string]ConsumeCallback
	stopWork  chan struct{}
}

// NewPulsarConsumer 创建PulsarConsumer.
func NewPulsarConsumer(conf *Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
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
		stopWork:  make(chan struct{}, 1),
	}

	topicVals := make([]string, 0)
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

func (pc *pulsarConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := pc.topics[topicKey]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}

	pc.callbacks[topicVal] = f

	return nil
}

func (pc *pulsarConsumer) Start() error {
	if len(pc.callbacks) == 0 {
		return errors.New(ErrSubscriberNil)
	}

	for {
		select {
		case <-pc.stopWork:
			return nil
		case cm := <-pc.channel:
			msg := cm.Message
			vals := strings.Split(msg.Topic(), "-partition-")
			if len(vals) != 2 {
				continue
			}
			fn, ok := pc.callbacks[vals[0]]
			if !ok {
				continue
			}

			data := &Message{
				Payload:    msg.Payload(),
				Topic:      vals[0],
				Properties: msg.Properties(),
			}
			if err := routine.Safe(func() {
				if err := fn(context.Background(), data); err == nil {
					_ = cm.Ack(msg)
				} else {
					cm.Nack(msg)
				}
			}); err != nil {
				// panic do not retry
				_ = cm.Ack(msg)
			}
		}
	}
}

func (pc *pulsarConsumer) Close() error {
	pc.stopWork <- struct{}{}
	pc.consumer.Close()
	pc.client.Close()

	return nil
}
