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
	breaker, ok = breakers[name]
	if !ok {
		breaker = newBreaker(name)
		breakers[name] = breaker
	}
	lock.Unlock()

	return breaker
}

func newBreaker(name string) Breaker {
	return newSREBreaker(name)
}
