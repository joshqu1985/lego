package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshqu1985/lego/metrics"
	"github.com/joshqu1985/lego/transport/naming"
)

type (
	// Router 避免业务代码直接引用
	Router  = gin.Engine
	Context = gin.Context
)

type Server struct {
	Name string
	Addr string

	naming naming.Naming
	router *Router
	option options
}

func NewServer(name, addr string, opts ...Option) (*Server, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		option.Naming = naming.NewPass(&naming.Config{})
	}
	if option.RouterRegister == nil {
		return nil, fmt.Errorf("router register function is nil")
	}

	server := &Server{
		Name:   name,
		Addr:   addr,
		naming: option.Naming,
		router: gin.New(),
		option: option,
	}
	server.router.Use(ServerRecover())
	server.router.Use(ServerCors())
	server.router.Use(ServerMetrics())

	{
		server.healthRegister()
		server.pprofRegister()
		server.metricRegister()
	}
	option.RouterRegister(server.router)

	return server, nil
}

func (this *Server) Start() error {
	if err := this.naming.Register(this.Name, this.Addr); err != nil {
		return err
	}

	httpsrv := &http.Server{Handler: this.router, Addr: this.Addr}

	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
	return nil
}

func (this *Server) AddMiddleware(inter gin.HandlerFunc) {
	this.router.Use(inter)
}

func (this *Server) healthRegister() {
	this.router.GET("/ping", func(ctx *Context) { ctx.String(200, "pong") })
}

func (this *Server) pprofRegister() {
	this.router.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	this.router.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	this.router.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	this.router.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	this.router.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	this.router.GET("/debug/pprof/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
	this.router.GET("/debug/pprof/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
	this.router.GET("/debug/pprof/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
	this.router.GET("/debug/pprof/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
	this.router.GET("/debug/pprof/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
}

func (this *Server) metricRegister() {
	if this.option.Metrics {
		metrics.ServeGIN(this.router)
	}
}
