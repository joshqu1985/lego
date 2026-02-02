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
		fmt.Println("-------------", time.Now().Unix(), string(msg.GetPayload()))
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

	msg3 := NewMessage([]byte("hello kitty 3"))
	err = producer.Send(context.Background(), "test", msg3)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	msg4 := NewMessage([]byte("hello kitty 4"))
	err = producer.Send(context.Background(), "test", msg4)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	time.Sleep(10 * time.Second)
}
