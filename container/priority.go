package container

import (
	"container/heap"
	"fmt"
	"sync"
)

// Priority 优先级队列
type Priority struct {
	Heap *Heap
	sync.RWMutex
}

// NewPriority 初始化优先级队列
func NewPriority() *Priority {
	pheap := &Heap{
		less: func(n1, n2 *heapNode) bool { return n1.Score < n2.Score },
	}
	heap.Init(pheap)

	return &Priority{Heap: pheap}
}

// Put 插入数据 score小的优先
func (this *Priority) Put(data any, score int64) {
	this.Lock()
	defer this.Unlock()

	heap.Push(this.Heap, &heapNode{Value: data, Score: score})
}

// Get 获取优先级最高的(score最小)数据
func (this *Priority) Get() (any, error) {
	this.Lock()
	defer this.Unlock()

	if this.Heap.Len() == 0 {
		return nil, fmt.Errorf("queue empty")
	}

	node := heap.Pop(this.Heap)
	data := node.(*heapNode)
	return data.Value, nil
}

func (this *Priority) Top() (any, int64, error) {
	this.RLock()
	defer this.RUnlock()

	if this.Heap.Len() == 0 {
		return nil, 0, fmt.Errorf("queue empty")
	}

	data := this.Heap.nodes[0]
	return data.Value, data.Score, nil
}

func (this *Priority) IsEmpty() bool {
	this.RLock()
	defer this.RUnlock()

	return this.Heap.Len() == 0
}

func (this *Priority) Size() int {
	this.RLock()
	defer this.RUnlock()

	return this.Heap.Len()
}

var _ heap.Interface = (*Heap)(nil)

type heapNode struct {
	Value any
	Score int64
}

type Heap struct {
	nodes []*heapNode
	less  func(self, other *heapNode) bool
}

func (this *Heap) Len() int {
	return len(this.nodes)
}

func (this *Heap) Less(i, j int) bool {
	return this.less(this.nodes[i], this.nodes[j])
}

func (this *Heap) Swap(i, j int) {
	this.nodes[i], this.nodes[j] = this.nodes[j], this.nodes[i]
}

func (this *Heap) Push(e any) {
	data := e.(*heapNode)
	this.nodes = append(this.nodes, data)
}

func (this *Heap) Pop() any {
	n := len(this.nodes)
	data := (this.nodes)[n-1]

	this.nodes = (this.nodes)[0 : n-1]
	return data
}
