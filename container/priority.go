package container

import (
	"container/heap"
	"sync"
)

type (
	// Priority 优先级队列.
	Priority struct {
		Heap *Heap
		sync.RWMutex
	}

	heapNode struct {
		Value any
		Score int64
	}

	Heap struct {
		less  func(self, other *heapNode) bool
		nodes []*heapNode
	}
)

var _ heap.Interface = (*Heap)(nil)

// NewPriority 初始化优先级队列.
func NewPriority() *Priority {
	pheap := &Heap{
		less: func(n1, n2 *heapNode) bool { return n1.Score < n2.Score },
	}
	heap.Init(pheap)

	return &Priority{Heap: pheap}
}

// Put 插入数据 score小的优先.
func (p *Priority) Put(data any, score int64) {
	p.Lock()
	defer p.Unlock()

	heap.Push(p.Heap, &heapNode{Value: data, Score: score})
}

// Get 获取优先级最高的(score最小)数据.
func (p *Priority) Get() (any, error) {
	p.Lock()
	defer p.Unlock()

	if p.Heap.Len() == 0 {
		return nil, ErrQueueEmpty
	}

	node := heap.Pop(p.Heap)
	data, _ := node.(*heapNode)

	return data.Value, nil
}

func (p *Priority) Top() (value any, score int64, err error) {
	p.RLock()
	defer p.RUnlock()

	value = nil
	score = 0
	if p.Heap.Len() == 0 {
		err = ErrQueueEmpty

		return
	}

	data := p.Heap.nodes[0]
	value, score = data.Value, data.Score

	return
}

func (p *Priority) IsEmpty() bool {
	p.RLock()
	defer p.RUnlock()

	return p.Heap.Len() == 0
}

func (p *Priority) Size() int {
	p.RLock()
	defer p.RUnlock()

	return p.Heap.Len()
}

func (p *Heap) Len() int {
	return len(p.nodes)
}

func (p *Heap) Less(i, j int) bool {
	return p.less(p.nodes[i], p.nodes[j])
}

func (p *Heap) Swap(i, j int) {
	p.nodes[i], p.nodes[j] = p.nodes[j], p.nodes[i]
}

func (p *Heap) Push(e any) {
	data, _ := e.(*heapNode)
	p.nodes = append(p.nodes, data)
}

func (p *Heap) Pop() any {
	n := len(p.nodes)
	data := p.nodes[n-1]

	p.nodes = p.nodes[:n-1]

	return data
}
