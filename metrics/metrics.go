package metrics

import (
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type (
	CounterOpt struct {
		Namespace string
		Name      string
		Help      string
		Labels    []string
	}

	CounterVec interface {
		Add(v float64, labels ...string)
		Inc(labels ...string)
	}

	GaugeOpt struct {
		Namespace string
		Name      string
		Help      string
		Labels    []string
	}

	GaugeVec interface {
		Set(v float64, labels ...string)
		Add(v float64, labels ...string)
		Sub(v float64, labels ...string)
	}

	HistogramOpt struct {
		Namespace string
		Name      string
		Help      string
		Labels    []string
		Buckets   []float64
	}

	HistogramVec interface {
		Observe(v int64, labels ...string)
		ObserveFloat(v float64, labels ...string)
	}
)

var enable int32

func NewCounter(opts *CounterOpt) CounterVec {
	return newPromeCounter(opts)
}

func NewGauge(opts *GaugeOpt) GaugeVec {
	return newPromeGauge(opts)
}

func NewHistogram(opts *HistogramOpt) HistogramVec {
	return newPromeHistogram(opts)
}

func Enabled() bool {
	return atomic.LoadInt32(&enable) == 1
}

func ServeHTTP() {
	prometheusHTTP()
	atomic.StoreInt32(&enable, 1)
}

func ServeGIN(router *gin.Engine) {
	prometheusGIN(router)
	atomic.StoreInt32(&enable, 1)
}
