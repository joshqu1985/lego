package container

import (
	"fmt"
	"sync"
)

// Node 链表节点
type Node struct {
	data any
	next *Node
}

// Linked 链表
type Linked struct {
	head *Node
	tail *Node
	size int
	lock sync.RWMutex
}

// NewLinked 初始化链表
func NewLinked() *Linked {
	return &Linked{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Get 获取队首元素
func (this *Linked) Get() (any, error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.size == 0 {
		return nil, fmt.Errorf("queue empty")
	}

	value := this.head.data
	this.head = this.head.next
	this.size--

	if this.size == 0 {
		this.tail = nil
	}
	return value, nil
}

// Put 插入数据到队列尾部
func (this *Linked) Put(e any) {
	this.lock.Lock()
	defer this.lock.Unlock()

	this.put(e)
}

func (this *Linked) Size() int {
	this.lock.RLock()
	defer this.lock.RUnlock()

	return this.size
}

func (this *Linked) IsEmpty() bool {
	this.lock.RLock()
	defer this.lock.RUnlock()

	return this.size == 0
}

func (this *Linked) put(e any) {
	node := &Node{data: e}

	if this.size == 0 {
		this.head = node
	} else {
		this.tail.next = node
	}

	this.tail = node
	this.size++
}
