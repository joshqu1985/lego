package rest

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshqu1985/lego/metrics"
)

func ServerCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if origin := ctx.Request.Header.Get("Origin"); origin != "" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,PATCH,DELETE,OPTIONS")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers",
				"Origin,X-Requested-With,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Accept,Authorization,AccessToken,Token")
			ctx.Writer.Header().Set("Access-Control-Expose-Headers",
				"Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Type,Content-Length,Content-Language")
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func ServerRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				reqbytes, _ := httputil.DumpRequest(ctx.Request, false)
				if isBrokenPipe(r) {
					log.Printf("[BROKEN-PIPE] time:%v request:%s err:%v\n", time.Now(), string(reqbytes), r)
					ctx.Error(errors.New("broken pipe"))
					ctx.Abort()
				} else {
					log.Printf("[PANIC] time:%v request:%s err:%v stack:%s\n", time.Now(), string(reqbytes), r, string(debug.Stack()))
					ctx.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		ctx.Next()
	}
}

var (
	serverMetricsDuration = metrics.NewHistogram(metrics.HistogramOpt{
		Namespace: "http_server",
		Name:      "duration_ms",
		Help:      "http server requests duration(ms).",
		Labels:    []string{"method"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	serverMetricsTotal = metrics.NewCounter(metrics.CounterOpt{
		Namespace: "http_server",
		Name:      "code_total",
		Help:      "http server requests total.",
		Labels:    []string{"method", "code"},
	})
)

func ServerMetrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		method := ctx.Request.Method + ctx.FullPath()
		serverMetricsDuration.Observe(time.Since(start).Milliseconds(), method)
		serverMetricsTotal.Inc(method, strconv.Itoa(ctx.Writer.Status()))
	}
}

func isBrokenPipe(r any) bool {
	var brokenPipe bool
	if opErr, ok := r.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(syscallErr.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(syscallErr.Error()), "connection reset by peer") {
				brokenPipe = true
			}
		}
	}
	return brokenPipe
}
