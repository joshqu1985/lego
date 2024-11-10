package broker

import (
	"context"
	"fmt"
	"os"
	"time"

	rocket "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

func NewRocketProducer(conf Config) (Producer, error) {
	topics := []string{}
	for _, val := range conf.Topics {
		topics = append(topics, val)
	}

	os.Setenv("mq.consoleAppender.enabled", "true")
	rocket.ResetLogger()
	client, err := rocket.NewProducer(&rocket.Config{
		Endpoint:      conf.Endpoints[0],
		ConsumerGroup: conf.GroupId,
		Credentials: &credentials.SessionCredentials{
			AccessKey: conf.AccessKey, AccessSecret: conf.SecretKey,
		},
		NameSpace: conf.AppId,
	},
		rocket.WithTopics(topics...),
	)
	if err != nil {
		return nil, err
	}

	return &rocketProducer{
		client: client,
		topics: conf.Topics,
	}, nil
}

type rocketProducer struct {
	client rocket.Producer
	topics map[string]string
}

func (this *rocketProducer) Send(ctx context.Context, topic string, msg *Message, args ...int64) error {
	if msg == nil {
		return fmt.Errorf("message is nil")
	}
	realTopic, ok := this.topics[topic]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	data := rocket.Message{
		Topic: realTopic,
		Body:  msg.Payload,
	}
	for key, val := range msg.Properties {
		if key == "tag" {
			data.SetTag(val)
		}
		data.AddProperty(key, val)
	}
	if len(args) != 0 && args[0] > time.Now().Unix() {
		data.SetDelayTimestamp(time.Unix(args[0], 0))
	}

	_, err := this.client.Send(ctx, &data)
	return err
}

func (this *rocketProducer) Close() error {
	return this.client.GracefulStop()
}
