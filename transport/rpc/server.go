package rpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"lego/transport/naming"
)

type Server struct {
	Name         string
	Addr         string
	registerFunc RegisterFunc

	register           naming.Naming
	options            []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

type ServerConfig struct {
	Name string `json:"name" yaml:"name" toml:"name"`
	Addr string `json:"addr" yaml:"addr" toml:"addr"`
}

func NewServer(conf ServerConfig, opts ...Option) (*Server, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		return nil, fmt.Errorf("naming is nil")
	}
	if option.RegisterFunc == nil {
		return nil, fmt.Errorf("register function is nil")
	}

	server := &Server{
		Name:         conf.Name,
		Addr:         conf.Addr,
		registerFunc: option.RegisterFunc,
		register:     option.Naming,
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

	this.registerFunc(server)

	if err := this.register.Register(this.Name, this.Addr); err != nil {
		return err
	}
	return server.Serve(lis)
}
