package rpc

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strconv"

	"google.golang.org/grpc"

	"github.com/joshqu1985/lego/metrics"
	"github.com/joshqu1985/lego/transport/naming"
)

type Server struct {
	Name string
	Addr string

	naming naming.Naming
	router RouterRegister
	option options

	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
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
		router: option.RouterRegister,
		option: option,
	}
	server.AddUnaryInterceptors(ServerUnaryMetrics())
	server.AddUnaryInterceptors(ServerUnaryRecover())
	server.AddStreamInterceptors(ServerStreamRecover())

	return server, nil
}

func (this *Server) Start() error {
	lis, err := net.Listen("tcp", this.Addr)
	if err != nil {
		return err
	}

	options := make([]grpc.ServerOption, 0)
	options = append(options,
		grpc.ChainUnaryInterceptor(this.unaryInterceptors...),
		grpc.ChainStreamInterceptor(this.streamInterceptors...))
	server := grpc.NewServer(options...)

	if err := this.naming.Register(this.Name, this.Addr); err != nil {
		return err
	}

	this.router(server)
	_ = this.httpServe(this.Addr)

	monitor := NewMonitor()
	monitor.AddShutdownCallback(func() {
		_ = this.naming.Deregister(this.Name)
		server.GracefulStop()
	})

	return server.Serve(lis)
}

func (this *Server) AddUnaryInterceptors(inter grpc.UnaryServerInterceptor) {
	this.unaryInterceptors = append(this.unaryInterceptors, inter)
}

func (this *Server) AddStreamInterceptors(inter grpc.StreamServerInterceptor) {
	this.streamInterceptors = append(this.streamInterceptors, inter)
}

func (this *Server) httpServe(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	iport, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("%s:%d", host, iport+1) // 端口+1

	if this.option.Metrics {
		metrics.ServeHTTP()
	}
	this.HealthServeHandle()

	go func() {
		_ = http.ListenAndServe(endpoint, nil)
	}()
	return nil
}

func (this *Server) HealthServeHandle() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}
