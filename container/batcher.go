package container

import (
	"sync"
	"time"
)

type Batcher struct {
	interval     time.Duration //
	bulkSize     int64         // 批量数据的最大条数
	entries      []any         //
	sync.RWMutex               //

	queueSize int64      //
	queue     chan []any // 存放批量数据的队列
}

func NewBatcher(opts ...BatcherOption) *Batcher {
	var batcher Batcher
	for _, opt := range opts {
		opt(&batcher)
	}
	if batcher.interval == 0 {
		batcher.interval = time.Second
	}
	if batcher.bulkSize == 0 {
		batcher.bulkSize = 500
	}
	if batcher.queueSize == 0 {
		batcher.queueSize = 100
	}
	batcher.entries = make([]any, 0, batcher.bulkSize)
	batcher.queue = make(chan []any, batcher.queueSize)

	go batcher.intervalFlush()
	return &batcher
}

func (this *Batcher) Put(data any) {
	this.Lock()
	defer this.Unlock()

	this.entries = append(this.entries, data)
	if len(this.entries) >= int(this.bulkSize) {
		this.flush()
	}
}

func (this *Batcher) Queue() chan []any {
	return this.queue
}

func (this *Batcher) intervalFlush() {
	ticker := time.NewTicker(this.interval)
	for {
		select {
		case <-ticker.C:
			this.Lock()
			this.flush()
			this.Unlock()
		}
	}
}

func (this *Batcher) flush() {
	if len(this.entries) == 0 {
		return
	}
	this.queue <- this.entries
	this.entries = make([]any, 0, this.bulkSize)
}

type BatcherOption func(o *Batcher)

// WithInterval 设置批量的最大等待时间
func WithInterval(interval time.Duration) BatcherOption {
	return func(b *Batcher) {
		b.interval = interval
	}
}

// WithBulkSize 设置批量的最大条数
func WithBulkSize(bulkSize int64) BatcherOption {
	return func(b *Batcher) {
		b.bulkSize = bulkSize
	}
}

// WithQueueSize 设置队列的最大长度
func WithQueueSize(queueSize int64) BatcherOption {
	return func(b *Batcher) {
		b.queueSize = queueSize
	}
}
