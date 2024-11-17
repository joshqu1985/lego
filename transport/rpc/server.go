package rpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/joshqu1985/lego/transport/naming"
)

type Server struct {
	Name string
	Addr string

	handlersRegister   HandlersRegister
	namingRegister     naming.Naming
	options            []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

type Config struct {
	Name string `json:"name" yaml:"name" toml:"name"`
	Addr string `json:"addr" yaml:"addr" toml:"addr"`
}

func NewServer(conf Config, opts ...Option) (*Server, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		option.Naming = naming.NewPass(&naming.Config{})
	}
	if option.HandlersRegister == nil {
		return nil, fmt.Errorf("handler register function is nil")
	}

	server := &Server{
		Name:             conf.Name,
		Addr:             conf.Addr,
		handlersRegister: option.HandlersRegister,
		namingRegister:   option.Naming,
	}
	return server, nil
}

func (this *Server) Start() error {
	lis, err := net.Listen("tcp", this.Addr)
	if err != nil {
		return err
	}

	options := append(this.options,
		grpc.ChainUnaryInterceptor(this.unaryInterceptors...),
		grpc.ChainStreamInterceptor(this.streamInterceptors...))
	server := grpc.NewServer(options...)

	this.handlersRegister(server)

	if err := this.namingRegister.Register(this.Name, this.Addr); err != nil {
		return err
	}

	monitor := NewMonitor()
	monitor.AddShutdownCallback(func() {
		_ = this.namingRegister.Deregister(this.Name)
		server.GracefulStop()
	})

	return server.Serve(lis)
}
