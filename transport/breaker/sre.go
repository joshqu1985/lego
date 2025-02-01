package breaker

import (
	"math"
	"sync"
	"time"

	"github.com/joshqu1985/lego/container"
	"golang.org/x/exp/rand"
)

type sreBreaker struct {
	name       string
	stat       *container.SlidingWindow[*statBucket]
	k          float64
	minRequest int64
	random     *rand.Rand
	randLock   sync.Mutex
}

func newSREBreaker(name string) Breaker {
	buckets := []*statBucket{}
	for i := 0; i < 20; i++ {
		buckets = append(buckets, &statBucket{})
	}

	return &sreBreaker{
		name:       name,
		k:          1.8,
		stat:       container.NewSlidingWindow(time.Duration(300)*time.Millisecond, buckets),
		minRequest: 100,
		random:     rand.New(rand.NewSource(uint64(time.Now().UnixNano()))),
	}
}

func (this *sreBreaker) Name() string {
	return this.name
}

func (this *sreBreaker) Allow() bool {
	total, access := this.summary()
	requests := this.k * float64(access)

	if total < this.minRequest || float64(total) < requests {
		return true
	}

	failRate := math.Max(0, (float64(total)-requests)/float64(total+1))
	return this.judgeAllow(failRate)
}

func (this *sreBreaker) MarkPass() {
	this.stat.Add(1)
}

func (this *sreBreaker) MarkFail() {
	this.stat.Add(0)
}

func (this *sreBreaker) summary() (int64, int64) {
	var (
		total, access int64
	)
	this.stat.Range(func(bucket *statBucket) {
		total += bucket.Total
		access += bucket.Access
	})
	return total, access
}

func (this *sreBreaker) judgeAllow(failRate float64) (allow bool) {
	this.randLock.Lock()
	allow = this.random.Float64() >= failRate
	this.randLock.Unlock()
	return
}

type statBucket struct {
	Access int64
	Total  int64
}

func (this *statBucket) Add(val int64) {
	this.Access += val
	this.Total++
}

func (this *statBucket) Reset() {
	this.Access = 0
	this.Total = 0
}
