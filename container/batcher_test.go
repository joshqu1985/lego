package container

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestBatcher(t *testing.T) {
	batcher := NewBatcher(WithBulkSize(3), WithInterval(time.Second))
	queue := batcher.Queue()
	go batcherTestConsume(queue)

	batcher.Put("hello")
	batcher.Put("hello")
	batcher.Put("hello")
	batcher.Put("hello")
	batcher.Put("hello")

	time.Sleep(3 * time.Second)
}

func BenchmarkBatcher(b *testing.B) {
	batcher := NewBatcher(WithBulkSize(10), WithInterval(time.Second))
	queue := batcher.Queue()
	go batcherTestConsume(queue)

	b.ResetTimer()

	for l := 0; l < b.N; l++ {
		batcher.Put(l)
	}
	b.StopTimer()
}

func batcherTestConsume(queue chan []any) {
	for {
		data := <-queue
		b, _ := json.Marshal(data)
		fmt.Println("len: ", len(data), " data: ", string(b))
	}
}
