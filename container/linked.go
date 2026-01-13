package container

import (
	"errors"
	"sync"
)

type (
	// Node 链表节点.
	Node struct {
		data any
		next *Node
	}

	// Linked 链表.
	Linked struct {
		head *Node
		tail *Node
		size int
		sync.RWMutex
	}
)

var ErrQueueEmpty = errors.New("queue empty")

// NewLinked 初始化链表.
func NewLinked() *Linked {
	return &Linked{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Get 获取队首元素.
func (lk *Linked) Get() (any, error) {
	lk.Lock()
	defer lk.Unlock()

	if lk.size == 0 {
		return nil, ErrQueueEmpty
	}

	value := lk.head.data
	lk.head = lk.head.next
	lk.size--

	if lk.size == 0 {
		lk.tail = nil
	}

	return value, nil
}

// Put 插入数据到队列尾部.
func (lk *Linked) Put(e any) {
	lk.Lock()
	defer lk.Unlock()

	lk.put(e)
}

func (lk *Linked) Size() int {
	lk.RLock()
	defer lk.RUnlock()

	return lk.size
}

func (lk *Linked) IsEmpty() bool {
	lk.RLock()
	defer lk.RUnlock()

	return lk.size == 0
}

func (lk *Linked) put(e any) {
	node := &Node{data: e}

	if lk.size == 0 {
		lk.head = node
	} else {
		lk.tail.next = node
	}

	lk.tail = node
	lk.size++
}
