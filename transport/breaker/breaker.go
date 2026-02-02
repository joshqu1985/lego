package breaker

import (
	"sync"
)

type Breaker interface {
	Allow() bool
	MarkPass()
	MarkFail()
}

var (
	lock     sync.RWMutex
	breakers = make(map[string]Breaker)
)

func Get(name string) Breaker {
	lock.RLock()
	breaker, ok := breakers[name]
	lock.RUnlock()
	if ok {
		return breaker
	}

	lock.Lock()
	defer lock.Unlock()

	if breaker, ok := breakers[name]; ok {
		return breaker
	}

	breaker = newBreaker(name)
	breakers[name] = breaker
	return breaker
}

func newBreaker(name string) Breaker {
	return newSREBreaker(name)
}
