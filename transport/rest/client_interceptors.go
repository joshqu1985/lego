package rest

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/joshqu1985/lego/metrics"
	"github.com/joshqu1985/lego/transport/breaker"
	"github.com/joshqu1985/lego/transport/naming"
)

func ClientResolver(target string, n naming.Naming) (string, error) {
	service := n.Service(target)
	addrs, err := service.Addrs()
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", fmt.Errorf("resolver target endpoint not found")
	}

	endpoint := addrs[rand.Intn(len(addrs))]
	if !strings.Contains(endpoint, "://") {
		endpoint = "http://" + endpoint
	}
	return endpoint, nil
}

func ClientBreakerAllow(method string) error {
	brk := breaker.Get(method)
	if !brk.Allow() {
		brk.MarkFail()
		return fmt.Errorf("circuit breaker not allowed")
	}
	return nil
}

func ClientBreakerMark(method string, code int) {
	brk := breaker.Get(method)
	if code == http.StatusInternalServerError ||
		code == http.StatusServiceUnavailable || code == http.StatusGatewayTimeout {
		brk.MarkFail()
	} else {
		brk.MarkPass()
	}
}

var (
	clientMetricsDuration = metrics.NewHistogram(metrics.HistogramOpt{
		Namespace: "http_client",
		Name:      "duration_ms",
		Help:      "http client requests duration(ms).",
		Labels:    []string{"method"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	clientMetricsTotal = metrics.NewCounter(metrics.CounterOpt{
		Namespace: "http_client",
		Name:      "code_total",
		Help:      "http client requests total.",
		Labels:    []string{"method", "code"},
	})
)

func ClientMetrics(method string, elapse int64, code int) {
	clientMetricsDuration.Observe(elapse, method)
	clientMetricsTotal.Inc(method, strconv.Itoa(code))
}
