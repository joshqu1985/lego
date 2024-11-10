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
func NewRocketConsumer(conf Config) (*rocketConsumer, error) {
	os.Setenv("mq.consoleAppender.enabled", "true")
	rocket.ResetLogger()
	client, err := rocket.NewSimpleConsumer(&rocket.Config{
		Endpoint: conf.Endpoints[0],
		Credentials: &credentials.SessionCredentials{
			AccessKey:    conf.AccessKey,
			AccessSecret: conf.SecretKey,
		},
		ConsumerGroup: conf.GroupId,
	},
		rocket.WithAwaitDuration(time.Second*5),
	)
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

func (this *rocketConsumer) Register(topic string, f ConsumeCallback) error {
	realTopic := ""
	if v, ok := this.topics[topic]; !ok {
		return fmt.Errorf("topic not found")
	} else {
		realTopic = v
	}

	this.callbacks[realTopic] = f
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
