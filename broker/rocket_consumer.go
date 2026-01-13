package broker

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5/credentials"

	rocket "github.com/apache/rocketmq-clients/golang/v5"

	"github.com/joshqu1985/lego/utils/routine"
)

const (
	MQConsoleAppenderEnabled = "mq.consoleAppender.enabled"
)

// rocketConsumer rocketmq消费者结构.
type rocketConsumer struct {
	client    rocket.SimpleConsumer
	topics    map[string]string
	callbacks map[string]ConsumeCallback
	stopWork  chan struct{}
}

// NewRocketConsumer 创建RocketConsumer.
func NewRocketConsumer(conf *Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
	}
	os.Setenv(MQConsoleAppenderEnabled, "true")
	rocket.ResetLogger()

	expr := make(map[string]*rocket.FilterExpression)
	for _, topicVal := range conf.Topics {
		expr[topicVal] = rocket.SUB_ALL
	}

	config := &rocket.Config{
		Endpoint:  conf.Endpoints[0],
		NameSpace: conf.AppId,
		Credentials: &credentials.SessionCredentials{
			AccessKey: conf.AccessKey, AccessSecret: conf.SecretKey,
		},
		ConsumerGroup: conf.GroupId,
	}
	client, err := rocket.NewSimpleConsumer(config, rocket.WithAwaitDuration(time.Second*5),
		rocket.WithSubscriptionExpressions(expr))
	if err != nil {
		return nil, err
	}

	consumer := &rocketConsumer{
		client:    client,
		topics:    conf.Topics,
		callbacks: make(map[string]ConsumeCallback),
		stopWork:  make(chan struct{}, 1),
	}

	return consumer, nil
}

func (rc *rocketConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := rc.topics[topicKey]
	if !ok {
		return errors.New(ErrTopicNotFound)
	}

	rc.callbacks[topicVal] = f

	return nil
}

func (rc *rocketConsumer) Start() error {
	if err := rc.client.Start(); err != nil {
		return err
	}

	for {
		select {
		case <-rc.stopWork:
			return nil
		default:
			msgs, err := rc.client.Receive(context.Background(), 5, time.Second*30)
			if err != nil {
				continue
			}
			for _, msg := range msgs {
				fn, ok := rc.callbacks[msg.GetTopic()]
				if !ok || fn == nil {
					_ = rc.client.Ack(context.Background(), msg)

					continue
				}

				data := &Message{
					Payload:    msg.GetBody(),
					Topic:      msg.GetTopic(),
					Properties: msg.GetProperties(),
				}
				if msg.GetTag() != nil {
					data.Properties["tag"] = *msg.GetTag()
				}

				if yerr := routine.Safe(func() {
					if xerr := fn(context.Background(), data); xerr == nil {
						_ = rc.client.Ack(context.Background(), msg)
					}
				}); yerr != nil {
					// panic do not retry
					_ = rc.client.Ack(context.Background(), msg)
				}
			}
		}
	}
}

func (rc *rocketConsumer) Close() error {
	rc.stopWork <- struct{}{}

	return rc.client.GracefulStop()
}
