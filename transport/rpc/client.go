package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gresolver "google.golang.org/grpc/resolver"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/resolver"
)

type Client struct {
	target string

	conn    *grpc.ClientConn
	builder gresolver.Builder
}

func NewClient(target string, opts ...Option) (*Client, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		option.Naming = naming.NewPass(&naming.Config{})
	}

	builder, err := resolver.New(option.Naming)
	if err != nil {
		return nil, err
	}
	client := &Client{
		target:  resolver.BuildTarget(option.Naming, target),
		builder: builder,
	}

	if err := client.dial(); err != nil {
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

func (this *Client) dial() error {
	conn, err := grpc.NewClient(this.target,
		grpc.WithResolvers(this.builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		return err
	}
	this.conn = conn
	return nil
}
