package broker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRedisBroker(t *testing.T) {
	consumer, _ := NewRedisConsumer(Config{
		Endpoints: []string{"127.0.0.1:6379"},
		GroupId:   "g1",
		Topics:    map[string]string{"test": "test"},
	})
	consumer.Register("test", func(ctx context.Context, msg *Message) error {
		fmt.Println(time.Now().Unix(), string(msg.Payload))
		return nil
	})
	go consumer.Start()

	producer, _ := NewRedisProducer(Config{
		Endpoints: []string{"127.0.0.1:6379"},
		Topics:    map[string]string{"test": "test"},
	})
	err := producer.Send(context.Background(), "test", &Message{
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

	curr := time.Now().Unix()
	err = producer.Send(context.Background(), "test", &Message{
		Payload: []byte("hello kitty 3"),
	}, curr+5)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Printf("curr time stamp:%d\n", curr)

	time.Sleep(10 * time.Second)
}

func BenchmarkRedisProducer(b *testing.B) {
	producer, _ := NewRedisProducer(Config{
		Endpoints: []string{"127.0.0.1:6379"},
		Topics:    map[string]string{"test": "test"},
	})

	ctx := context.Background()

	b.Run("redis producer", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i <= b.N; i++ {
			_ = producer.Send(ctx, "test", &Message{
				Payload: []byte("hello kitty"),
			})
		}
	})
}
