package broker

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/glog"
	"github.com/pkg/errors"
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
		go this.work(topic, f)
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
			if err := this.process(&msg, f); err != nil {
				glog.Errorf("mem queue process err:%v", err)
				continue
			}
		case <-this.stopWork:
			return
		}
	}
}

func (this *memoryConsumer) process(msg *Message, f ConsumeCallback) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = errors.Wrap(e, "consumer panic")
				return
			}
			err = errors.Errorf("consumer panic: %v", r)
		}
	}()
	err = f(context.Background(), msg)
	return
}
