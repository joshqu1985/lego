package rpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"

	_ "net/http/pprof" //nolint:gosec

	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/metrics"
	"github.com/joshqu1985/lego/transport/naming"
)

type Server struct {
	Name string
	Addr string

	naming  naming.Naming
	grpcSrv *grpc.Server
	httpSrv *http.Server
	router  RouterRegister
	option  options

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
		return nil, errors.New("router register function is nil")
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

func (s *Server) Start() error {
	var lc net.ListenConfig
	lis, err := lc.Listen(context.Background(), "tcp", s.Addr)
	if err != nil {
		return err
	}

	options := make([]grpc.ServerOption, 0)
	options = append(options,
		grpc.ChainUnaryInterceptor(s.unaryInterceptors...),
		grpc.ChainStreamInterceptor(s.streamInterceptors...))
	s.grpcSrv = grpc.NewServer(options...)

	if xerr := s.naming.Register(s.Name, s.Addr); xerr != nil {
		return xerr
	}

	s.router(s.grpcSrv)
	if xerr := s.httpServe(s.Addr); xerr != nil {
		logs.Warnf("health check http serve err:%v", xerr)
	}

	monitor := NewMonitor()
	monitor.AddShutdownCallback(func() {
		_ = s.naming.Deregister(s.Name)
		s.grpcSrv.GracefulStop()
	})

	return s.grpcSrv.Serve(lis)
}

func (s *Server) AddUnaryInterceptors(inter grpc.UnaryServerInterceptor) {
	s.unaryInterceptors = append(s.unaryInterceptors, inter)
}

func (s *Server) AddStreamInterceptors(inter grpc.StreamServerInterceptor) {
	s.streamInterceptors = append(s.streamInterceptors, inter)
}

func (s *Server) httpServe(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	iport, err := strconv.ParseUint(port, 10, 64)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("%s:%d", host, iport+1) // 端口+1

	if s.option.Metrics {
		metrics.ServeHTTP()
	}
	s.HealthServeHandle()

	s.httpSrv = &http.Server{
		Addr:              endpoint,
		ReadHeaderTimeout: time.Second,
	}
	go func() {
		if xerr := s.httpSrv.ListenAndServe(); xerr != nil && xerr != http.ErrServerClosed {
			logs.Error("health server err:%v", xerr)
		}
	}()

	return nil
}

func (s *Server) HealthServeHandle() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
}
