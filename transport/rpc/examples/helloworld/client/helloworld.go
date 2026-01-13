package main

import (
	"context"
	"log"
	"time"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
)

type Helloworld struct {
	client pb.GreeterClient
}

func NewHelloworld(n naming.Naming) *Helloworld {
	c, err := rpc.NewClient("helloworld", rpc.WithNaming(n))
	if err != nil {
		log.Printf("rpc.NewClient failed: %v", err)

		return nil
	}

	return &Helloworld{client: pb.NewGreeterClient(c.Conn())}
}

func (h *Helloworld) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	xctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := h.client.SayHello(xctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Printf("could not greet: %v", err)

		return resp, err
	}

	return resp, nil
}
