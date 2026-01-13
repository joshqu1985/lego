package broker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPulsarBroker(t *testing.T) {
	consumer, err := NewPulsarConsumer(&Config{
		Endpoints: []string{"pulsar://localhost:6650"},
		GroupId:   "g1",
		Topics:    map[string]string{"test": "persistent://public/default/test"},
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

	producer, err := NewPulsarProducer(&Config{
		Endpoints: []string{"pulsar://localhost:6650"},
		Topics:    map[string]string{"test": "persistent://public/default/test"},
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = producer.Send(context.Background(), "test", &Message{
		Payload: []byte("hello kitty 3"),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = producer.Send(context.Background(), "test", &Message{
		Payload: []byte("hello kitty 4"),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	time.Sleep(10 * time.Second)
}
