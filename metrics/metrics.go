package metrics

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type CounterOpt struct {
	Namespace string
	Name      string
	Help      string
	Labels    []string
}

type CounterVec interface {
	Add(v float64, labels ...string)
	Inc(labels ...string)
}

func NewCounter(opts CounterOpt) CounterVec {
	return newPromeCounter(opts)
}

type GaugeOpt struct {
	Namespace string
	Name      string
	Help      string
	Labels    []string
}

type GaugeVec interface {
	Set(v float64, labels ...string)
	Add(v float64, labels ...string)
	Sub(v float64, labels ...string)
}

func NewGauge(opts GaugeOpt) GaugeVec {
	return newPromeGauge(opts)
}

type HistogramOpt struct {
	Namespace string
	Name      string
	Help      string
	Labels    []string
	Buckets   []float64
}

type HistogramVec interface {
	Observe(v int64, labels ...string)
	ObserveFloat(v float64, labels ...string)
}

func NewHistogram(opts HistogramOpt) HistogramVec {
	return newPromeHistogram(opts)
}

var (
	enable int32
)

func Enabled() bool {
	return atomic.LoadInt32(&enable) == 1
}

func ServeHandle() {
	http.Handle("/metrics", promhttp.Handler())
	atomic.StoreInt32(&enable, 1)
}

func ServeGin(router *gin.Engine) {
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	atomic.StoreInt32(&enable, 1)
}
