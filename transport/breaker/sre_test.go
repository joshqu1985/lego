package breaker

import (
	"fmt"
	"testing"
)

func TestSREAllow(t *testing.T) {
	brk := newSREBreaker("test")
	markPass(brk, 200)
	if !brk.Allow() {
		t.Fatal("allow should be true")
	}
}

func TestSREHalfOpen(t *testing.T) {
	brk := newSREBreaker("test")
	markPass(brk, 100)
	markFail(brk, 3000)
	fmt.Println(brk.Allow())
}

func BenchmarkSreBreakerAllow(b *testing.B) {
	brk := newSREBreaker("test")
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		_ = brk.Allow()
		if i%2 == 0 {
			brk.MarkPass()
		} else {
			brk.MarkFail()
		}
	}
}

func markPass(b Breaker, count int) {
	for i := 0; i < count; i++ {
		b.MarkPass()
	}
}

func markFail(b Breaker, count int) {
	for i := 0; i < count; i++ {
		b.MarkFail()
	}
}
