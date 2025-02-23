package broker

import (
	"context"
	"fmt"
	"os"
	"time"

	rocket "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

// NewRocketConsumer 创建RocketConsumer
func NewRocketConsumer(conf Config) (Consumer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}
	os.Setenv("mq.consoleAppender.enabled", "true")
	rocket.ResetLogger()

	expr := map[string]*rocket.FilterExpression{}
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
	}
	return consumer, nil
}

// rocketConsumer rocketmq消费者结构
type rocketConsumer struct {
	client    rocket.SimpleConsumer
	topics    map[string]string
	callbacks map[string]ConsumeCallback
}

func (this *rocketConsumer) Register(topicKey string, f ConsumeCallback) error {
	topicVal, ok := this.topics[topicKey]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	this.callbacks[topicVal] = f
	return nil
}

func (this *rocketConsumer) Start() error {
	if err := this.client.Start(); err != nil {
		return err
	}

	for {
		msgs, err := this.client.Receive(context.Background(), 5, time.Second*30)
		if err != nil {
			continue
		}

		for _, msg := range msgs {
			callback, ok := this.callbacks[msg.GetTopic()]
			if !ok || callback == nil {
				this.client.Ack(context.Background(), msg)
				continue
			}

			data := &Message{
				Topic:      msg.GetTopic(),
				Properties: msg.GetProperties(),
				Payload:    msg.GetBody(),
			}
			if msg.GetTag() != nil {
				data.Properties["tag"] = *(msg.GetTag())
			}
			if err := callback(context.Background(), data); err == nil {
				this.client.Ack(context.Background(), msg)
			}
		}
	}
}

func (this *rocketConsumer) Close() error {
	return this.client.GracefulStop()
}
