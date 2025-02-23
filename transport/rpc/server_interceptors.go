package rpc

import (
	"context"
	"log"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/joshqu1985/lego/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServerUnaryRecover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any,
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = panicToError(ctx, r)
			}
		}()
		return handler(ctx, req)
	}
}

func ServerStreamRecover() grpc.StreamServerInterceptor {
	return func(srv any, stream grpc.ServerStream,
		info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = panicToError(context.Background(), r)
			}
		}()
		return handler(srv, stream)
	}
}

var (
	serverMetricsDuration = metrics.NewHistogram(metrics.HistogramOpt{
		Namespace: "rpc_server",
		Name:      "duration_ms",
		Labels:    []string{"method"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	serverMetricsTotal = metrics.NewCounter(metrics.CounterOpt{
		Namespace: "rpc_server",
		Name:      "code_total",
		Labels:    []string{"method", "code"},
	})
)

func ServerUnaryMetrics() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any,
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		serverMetricsDuration.Observe(time.Since(start).Milliseconds(), info.FullMethod)
		serverMetricsTotal.Inc(info.FullMethod, strconv.Itoa(int(status.Code(err))))
		return resp, err
	}
}

func panicToError(_ context.Context, r any) error {
	log.Printf("[PANIC] time:%v err:%v stack:%s\n", time.Now(), r, string(debug.Stack()))

	if err, ok := r.(error); ok {
		return err
	}
	return status.Errorf(codes.Unknown, "%v", r)
}
