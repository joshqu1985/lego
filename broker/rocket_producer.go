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
	if len(conf.Endpoints) == 0 {
		return nil, fmt.Errorf("endpoints is empty")
	}
	os.Setenv("mq.consoleAppender.enabled", "true")
	rocket.ResetLogger()

	topicVals := []string{}
	for _, topicVal := range conf.Topics {
		topicVals = append(topicVals, topicVal)
	}

	config := &rocket.Config{
		Endpoint:      conf.Endpoints[0],
		ConsumerGroup: conf.GroupId,
		Credentials: &credentials.SessionCredentials{
			AccessKey: conf.AccessKey, AccessSecret: conf.SecretKey,
		},
		NameSpace: conf.AppId,
	}
	client, err := rocket.NewProducer(config, rocket.WithTopics(topicVals...))
	if err != nil {
		return nil, err
	}

	if err := client.Start(); err != nil {
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

	topicVal, ok := this.topics[topic]
	if !ok {
		return fmt.Errorf("topic not found")
	}

	data := rocket.Message{
		Topic: topicVal,
		Body:  msg.Payload,
	}
	if msg.Key != "" {
		data.SetKeys(msg.Key)
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

	receives, err := this.client.Send(ctx, &data)
	if len(receives) > 0 {
		msg.MessageId = receives[0].MessageID
	}
	return err
}

func (this *rocketProducer) Close() error {
	return this.client.GracefulStop()
}
