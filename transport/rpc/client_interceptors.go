package rpc

import (
	"context"
	"errors"
	"path"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/joshqu1985/lego/metrics"
	"github.com/joshqu1985/lego/transport/breaker"
)

var (
	clientMetricsDuration = metrics.NewHistogram(&metrics.HistogramOpt{
		Namespace: "rpc_client",
		Name:      "exec_duration",
		Help:      "rpc client exec duration",
		Labels:    []string{"method"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	clientMetricsTotal = metrics.NewCounter(&metrics.CounterOpt{
		Namespace: "rpc_client",
		Name:      "code_total",
		Help:      "rpc client request total",
		Labels:    []string{"method", "code"},
	})
)

func ClientUnaryTimeout(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any,
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		timedCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return invoker(timedCtx, method, req, reply, cc, opts...)
	}
}

func ClientUnaryBreaker() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any,
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		brk := breaker.Get(path.Join(cc.Target(), method))
		if !brk.Allow() {
			brk.MarkFail()

			return errors.New("circuit breaker not allowed")
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil && isErrMarkFail(err) {
			brk.MarkFail()
		} else {
			brk.MarkPass()
		}

		return err
	}
}

func isErrMarkFail(err error) bool {
	switch status.Code(err) {
	case codes.DeadlineExceeded, codes.Internal, codes.Unavailable, codes.DataLoss,
		codes.Unimplemented, codes.ResourceExhausted:
		return true
	default:
		return false
	}
}

func ClientUnaryMetrics() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any,
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		now := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		clientMetricsDuration.Observe(time.Since(now).Milliseconds(), method)
		clientMetricsTotal.Inc(method, strconv.Itoa(int(status.Code(err))))

		return err
	}
}
