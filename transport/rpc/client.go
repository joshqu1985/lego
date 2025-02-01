package rpc

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gresolver "google.golang.org/grpc/resolver"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc/resolver"
)

type Client struct {
	target  string
	conn    *grpc.ClientConn
	builder gresolver.Builder

	streamInterceptors []grpc.StreamClientInterceptor
	unaryInterceptors  []grpc.UnaryClientInterceptor
}

func NewClient(target string, opts ...Option) (*Client, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		option.Naming = naming.NewPass(&naming.Config{})
	}
	if option.Timeout == 0 {
		option.Timeout = time.Duration(3) * time.Second
	}

	builder, err := resolver.New(option.Naming)
	if err != nil {
		return nil, err
	}
	client := &Client{
		target:  resolver.BuildTarget(option.Naming, target),
		builder: builder,
	}
	client.addUnaryInterceptors(ClientUnaryTimeout(option.Timeout))
	client.addUnaryInterceptors(ClientUnaryBreaker())
	client.addUnaryInterceptors(ClientUnaryMetrics())

	options := []grpc.DialOption{
		grpc.WithResolvers(client.builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}
	options = append(options, grpc.WithChainUnaryInterceptor(client.unaryInterceptors...))
	options = append(options, grpc.WithChainStreamInterceptor(client.streamInterceptors...))

	if err := client.dial(options); err != nil {
		return nil, err
	}
	return client, nil
}

func (this *Client) Conn() *grpc.ClientConn {
	return this.conn
}

func (this *Client) Target() string {
	return this.target
}

func (this *Client) dial(options []grpc.DialOption) error {
	conn, err := grpc.NewClient(this.target, options...)
	if err != nil {
		return err
	}
	this.conn = conn
	return nil
}

func (this *Client) addUnaryInterceptors(inter grpc.UnaryClientInterceptor) {
	this.unaryInterceptors = append(this.unaryInterceptors, inter)
}

func (this *Client) addStreamInterceptors(inter grpc.StreamClientInterceptor) {
	this.streamInterceptors = append(this.streamInterceptors, inter)
}
