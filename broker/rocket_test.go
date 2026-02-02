package broker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRocketBroker(t *testing.T) {
	consumer, err := NewRocketConsumer(&Config{
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
		fmt.Println("-------------", time.Now().Unix(), string(msg.GetPayload()))
		return nil
	})
	go consumer.Start()

	producer, err := NewRocketProducer(&Config{
		Endpoints: []string{"localhost:8081"},
		Topics:    map[string]string{"test": "test"},
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	msg1 := NewMessage([]byte("hello kitty 1"))
	err = producer.Send(context.Background(), "test", msg1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	msg2 := NewMessage([]byte("hello kitty 2"))
	err = producer.Send(context.Background(), "test", msg2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	time.Sleep(10 * time.Second)
}
