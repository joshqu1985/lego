package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type promeCounter struct {
	counter *prometheus.CounterVec
}

func newPromeCounter(opts CounterOpt) CounterVec {
	vec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
	}, opts.Labels)
	prometheus.MustRegister(vec)

	return &promeCounter{counter: vec}
}

func (p *promeCounter) Add(v float64, labels ...string) {
	if Enabled() {
		p.counter.WithLabelValues(labels...).Add(v)
	}
}

func (p *promeCounter) Inc(labels ...string) {
	if Enabled() {
		p.counter.WithLabelValues(labels...).Inc()
	}
}

type promeGauge struct {
	gauge *prometheus.GaugeVec
}

func newPromeGauge(opts GaugeOpt) GaugeVec {
	vec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
	}, opts.Labels)
	prometheus.MustRegister(vec)

	return &promeGauge{gauge: vec}
}

func (p *promeGauge) Set(v float64, labels ...string) {
	if Enabled() {
		p.gauge.WithLabelValues(labels...).Set(v)
	}
}

func (p *promeGauge) Add(v float64, labels ...string) {
	if Enabled() {
		p.gauge.WithLabelValues(labels...).Add(v)
	}
}

func (p *promeGauge) Sub(v float64, labels ...string) {
	if Enabled() {
		p.gauge.WithLabelValues(labels...).Sub(v)
	}
}

type promeHistogram struct {
	histogram *prometheus.HistogramVec
}

func newPromeHistogram(opts HistogramOpt) HistogramVec {
	vec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
		Buckets:   opts.Buckets,
	}, opts.Labels)
	prometheus.MustRegister(vec)

	return &promeHistogram{histogram: vec}
}

func (p *promeHistogram) Observe(v int64, labels ...string) {
	if Enabled() {
		p.histogram.WithLabelValues(labels...).Observe(float64(v))
	}
}

func (p *promeHistogram) ObserveFloat(v float64, labels ...string) {
	if Enabled() {
		p.histogram.WithLabelValues(labels...).Observe(v)
	}
}
