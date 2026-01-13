package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/metrics"
	"github.com/joshqu1985/lego/transport/naming"
)

type (
	// Router 避免业务代码直接引用.
	Router  = gin.Engine
	Context = gin.Context

	Server struct {
		naming   naming.Naming
		httpsrv  *http.Server
		router   *Router
		quitChan chan os.Signal
		Name     string
		Addr     string
		option   options
		lock     sync.RWMutex
	}
)

var (
	StringBrokenPipe      = "broken pipe"
	StringConnectionReset = "connection reset by peer"
)

func NewServer(name, addr string, opts ...Option) (*Server, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		option.Naming = naming.NewPass(&naming.Config{})
	}
	if option.RouterRegister == nil {
		return nil, errors.New("router register function is nil")
	}

	server := &Server{
		Name:     name,
		Addr:     addr,
		naming:   option.Naming,
		router:   gin.New(),
		option:   option,
		quitChan: make(chan os.Signal, 1),
	}
	signal.Notify(server.quitChan, syscall.SIGINT, syscall.SIGTERM)

	server.router.Use(ServerRecover())
	server.router.Use(ServerCors())
	server.router.Use(ServerMetrics())

	server.healthRegister()
	server.pprofRegister()
	server.metricRegister()
	option.RouterRegister(server.router)

	return server, nil
}

func (s *Server) Start() error {
	if err := s.naming.Register(s.Name, s.Addr); err != nil {
		return err
	}

	s.lock.Lock()
	s.httpsrv = &http.Server{
		Handler:           s.router,
		Addr:              s.Addr,
		ReadHeaderTimeout: time.Second,
	}
	s.lock.Unlock()

	errChan := make(chan error, 1)
	go func() {
		if err := s.httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("http server err:%w", err)
		}
	}()

	// 等待信号或错误
	select {
	case err := <-errChan:
		// 处理服务错误
		_ = s.naming.Deregister(s.Name)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = s.httpsrv.Shutdown(ctx)

		return err
	case <-s.quitChan:
		// 处理优雅关闭
		logs.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.httpsrv.Shutdown(ctx); err != nil {
			logs.Errorf("Server shutdown error: %v\n", err)
		}
		_ = s.naming.Deregister(s.Name)
		logs.Info("Server exited")

		return nil
	}
}

func (s *Server) Close() error {
	select {
	case s.quitChan <- syscall.SIGTERM:
		// 信号已发送
	default:
		// 通道已满，避免阻塞
	}

	return nil
}

func (s *Server) AddMiddleware(inter gin.HandlerFunc) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.router.Use(inter)
}

func (s *Server) healthRegister() {
	s.router.GET("/ping", func(ctx *Context) { ctx.String(200, "pong") })
}

func (s *Server) pprofRegister() {
	s.router.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	s.router.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	s.router.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	s.router.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	s.router.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	s.router.GET("/debug/pprof/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
	s.router.GET("/debug/pprof/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
	s.router.GET("/debug/pprof/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
	s.router.GET("/debug/pprof/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
	s.router.GET("/debug/pprof/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
}

func (s *Server) metricRegister() {
	if s.option.Metrics {
		metrics.ServeGIN(s.router)
	}
}
