package container

import (
	"fmt"
	"testing"
	"time"
)

type TestBucket struct {
	Value int64
	Count int64
}

func (this *TestBucket) Add(cmd int64) {
	this.Value += cmd
	this.Count++
}

func (this *TestBucket) Reset() {
	this.Value = 0
	this.Count = 0
}

func TestSlidingWindowAdd(t *testing.T) {
	buckets := []*TestBucket{}
	for i := 0; i < 3; i++ {
		buckets = append(buckets, &TestBucket{})
	}

	sw := NewSlidingWindow(time.Duration(50)*time.Millisecond, buckets)

	sw.Add(1)
	time.Sleep(time.Duration(50) * time.Millisecond)
	sw.Add(2)
	sw.Add(3)
	time.Sleep(time.Duration(50) * time.Millisecond)
	sw.Add(4)
	sw.Add(5)
	sw.Add(6)
	time.Sleep(time.Duration(200) * time.Millisecond)
	sw.Add(7)
	sw.Add(8)
	sw.Add(9)

	sw.Range(func(bucket *TestBucket) {
		fmt.Println("---", bucket.Count, bucket.Value)
	})
}
