package broker

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/glog"
	"github.com/joshqu1985/lego/utils/routine"
)

func NewMemoryConsumer(conf Config) (Consumer, error) {
	return &memoryConsumer{
		consumers: map[string]ConsumeCallback{},
		groupId:   conf.GroupId,
		broker:    getMemoryBroker(conf.Topics),
		wg:        &sync.WaitGroup{},
		stopWork:  make(chan struct{}, len(conf.Topics)),
	}, nil
}

type memoryConsumer struct {
	consumers map[string]ConsumeCallback
	groupId   string
	broker    *memoryBroker
	wg        *sync.WaitGroup
	stopWork  chan struct{}
}

func (this *memoryConsumer) Register(topic string, f ConsumeCallback) error {
	this.consumers[topic] = f
	this.broker.addSubscriber(topic, this.groupId)
	return nil
}

func (this *memoryConsumer) Start() error {
	if len(this.consumers) == 0 {
		return fmt.Errorf("at least one consumer function registered")
	}

	for topic, f := range this.consumers {
		this.wg.Add(1)
		routine.Go(func() { this.work(topic, f) })
	}
	this.wg.Wait()

	return nil
}

func (this *memoryConsumer) Close() error {
	for i := 0; i < len(this.consumers); i++ {
		this.stopWork <- struct{}{}
	}

	for topic := range this.consumers {
		this.broker.delSubscriber(topic, this.groupId)
	}
	return nil
}

func (this *memoryConsumer) work(topic string, f ConsumeCallback) {
	defer this.wg.Done()

	recv, err := this.broker.getSubChannel(topic, this.groupId)
	if err != nil {
		return
	}

	for {
		select {
		case msg := <-recv:
			routine.Safe(func() {
				if err := f(context.Background(), &msg); err != nil {
					glog.Errorf("memory queue process err:%v", err)
				}
			})
		case <-this.stopWork:
			return
		}
	}
}
