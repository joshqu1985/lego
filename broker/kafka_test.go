package broker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestKafkaBroker(t *testing.T) {
	consumer, _ := NewKafkaConsumer(&Config{
		Endpoints: []string{"127.0.0.1:9092"},
		GroupId:   "g1",
		Topics:    map[string]string{"test": "test"},
	})
	consumer.Register("test", func(ctx context.Context, msg *Message) error {
		fmt.Println(time.Now().Unix(), string(msg.GetPayload()))
		return nil
	})
	go consumer.Start()

	producer, _ := NewKafkaProducer(&Config{
		Endpoints: []string{"127.0.0.1:9092"},
		Topics:    map[string]string{"test": "test"},
	})

	msg1 := NewMessage([]byte("hello kitty 1"))
	err := producer.Send(context.Background(), "test", msg1)
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

	curr := time.Now().Unix()
	msg3 := NewMessage([]byte("hello kitty 3"))
	err = producer.Send(context.Background(), "test", msg3)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Printf("curr time stamp:%d\n", curr)

	time.Sleep(10 * time.Second)
}

func BenchmarkKafkaProducer(b *testing.B) {
	producer, _ := NewKafkaProducer(&Config{
		Endpoints: []string{"127.0.0.1:9092"},
		Topics:    map[string]string{"test": "test"},
	})

	ctx := context.Background()

	b.Run("redis producer", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i <= b.N; i++ {
			msgx := NewMessage([]byte("hello kitty"))
			_ = producer.Send(ctx, "test", msgx)
		}
	})
}
