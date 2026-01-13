package breaker

import (
	"math"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/joshqu1985/lego/container"
)

type (
	sreBreaker struct {
		stat       *container.SlidingWindow[*statBucket]
		name       string
		k          float64
		minRequest int64
		lock       sync.Mutex
	}

	statBucket struct {
		Access int64
		Total  int64
	}
)

func newSREBreaker(name string) Breaker {
	buckets := make([]*statBucket, 0)
	for range 20 {
		buckets = append(buckets, &statBucket{})
	}

	return &sreBreaker{
		name:       name,
		k:          1.8,
		stat:       container.NewSlidingWindow(time.Duration(300)*time.Millisecond, buckets),
		minRequest: 100,
	}
}

func (srb *sreBreaker) Name() string {
	return srb.name
}

func (srb *sreBreaker) Allow() bool {
	total, access := srb.summary()
	requests := srb.k * float64(access)

	if total < srb.minRequest || float64(total) < requests {
		return true
	}

	failRate := math.Max(0, (float64(total)-requests)/float64(total+1))

	return srb.judgeAllow(failRate)
}

func (srb *sreBreaker) MarkPass() {
	srb.stat.Add(1)
}

func (srb *sreBreaker) MarkFail() {
	srb.stat.Add(0)
}

func (srb *sreBreaker) summary() (int64, int64) {
	var (
		total  int64
		access int64
	)
	srb.stat.Range(func(bucket *statBucket) {
		total += bucket.Total
		access += bucket.Access
	})

	return total, access
}

func (srb *sreBreaker) judgeAllow(failRate float64) bool {
	var allow bool
	srb.lock.Lock()
	allow = rand.Float64() >= failRate //nolint:gosec
	srb.lock.Unlock()

	return allow
}

func (srb *statBucket) Add(val int64) {
	srb.Access += val
	srb.Total++
}

func (srb *statBucket) Reset() {
	srb.Access = 0
	srb.Total = 0
}
