package container

import (
	"sync"
	"time"
)

type SlidingWindow[T IBucket] struct {
	buckets  []T
	size     int64
	offset   int64
	slide    time.Duration
	lastTime time.Time // start time of the last bucket
	sync.RWMutex
}

func NewSlidingWindow[T IBucket](slide time.Duration, buckets []T) *SlidingWindow[T] {
	return &SlidingWindow[T]{
		buckets:  buckets,
		size:     int64(len(buckets)),
		slide:    slide,
		lastTime: time.Now(),
	}
}

func (this *SlidingWindow[T]) Add(cmd int64) {
	this.Lock()
	defer this.Unlock()
	this.sliding()
	this.add(this.offset, cmd)
}

func (this *SlidingWindow[T]) Range(fn func(T)) {
	this.Lock()
	defer this.Unlock()
	this.sliding()
	for i := int64(0); i < this.size; i++ {
		fn(this.buckets[(i+this.offset+1)%this.size])
	}
}

func (this *SlidingWindow[T]) Size() int64 {
	return this.size
}

func (this *SlidingWindow[T]) sliding() {
	timespan := this.span()
	if timespan == 0 {
		return
	}

	offset := this.offset
	for i := int64(0); i < timespan; i++ {
		this.reset((offset + i + 1) % this.size)
	}

	this.offset = (offset + timespan) % this.size
	this.lastTime = this.lastTime.Add(time.Duration(timespan) * this.slide)
}

func (this *SlidingWindow[T]) span() int64 {
	return int64(time.Since(this.lastTime) / this.slide)
}

func (this *SlidingWindow[T]) add(offset int64, cmd int64) {
	this.buckets[offset%this.size].Add(cmd)
}

func (this *SlidingWindow[T]) reset(offset int64) {
	this.buckets[offset%this.size].Reset()
}

type IBucket interface {
	Add(int64)
	Reset()
}
