package broker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRocketBroker(t *testing.T) {
	consumer, err := NewRocketConsumer(Config{
		Endpoints: []string{"localhost:8081"},
		GroupId:   "g1",
		Topics:    map[string]string{"test": "test"},
		AppId:     "",
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	consumer.Register("test", func(ctx context.Context, msg *Message) error {
		fmt.Println("-------------", time.Now().Unix(), string(msg.Payload))
		return nil
	})
	go consumer.Start()

	producer, err := NewRocketProducer(Config{
		Endpoints: []string{"localhost:8081"},
		Topics:    map[string]string{"test": "test"},
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = producer.Send(context.Background(), "test", &Message{
		Payload: []byte("hello kitty 1"),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = producer.Send(context.Background(), "test", &Message{
		Payload: []byte("hello kitty 2"),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	time.Sleep(10 * time.Second)
}
