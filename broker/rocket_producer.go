package broker

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5/credentials"

	rocket "github.com/apache/rocketmq-clients/golang/v5"
)

type rocketProducer struct {
	client rocket.Producer
	topics map[string]string
}

func NewRocketProducer(conf *Config) (Producer, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New(ErrEndpointsEmpty)
	}
	os.Setenv(MQConsoleAppenderEnabled, "true")
	rocket.ResetLogger()

	topicVals := make([]string, 0)
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

	if xerr := client.Start(); xerr != nil {
		return nil, xerr
	}

	return &rocketProducer{
		client: client,
		topics: conf.Topics,
	}, nil
}

func (rp *rocketProducer) Send(ctx context.Context, topic string, msg *Message) error {
	return rp.SendDelay(ctx, topic, msg, 0)
}

func (rp *rocketProducer) SendDelay(ctx context.Context, topic string, msg *Message, stamp int64) error {
	if msg == nil {
		return errors.New(ErrMessageIsNil)
	}

	topicVal, ok := rp.topics[topic]
	if !ok {
		return errors.New(ErrTopicNotFound)
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

	if stamp > time.Now().Unix() {
		data.SetDelayTimestamp(time.Unix(stamp, 0))
	}

	receives, err := rp.client.Send(ctx, &data)
	if len(receives) > 0 {
		msg.MessageId = receives[0].MessageID
	}

	return err
}

func (rp *rocketProducer) Close() error {
	return rp.client.GracefulStop()
}
