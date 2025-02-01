package main

import (
	"context"
	"log"
	"time"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
)

func NewHelloworld() *Helloworld {
	c, err := rpc.NewClient("helloworld", rpc.WithNaming(naming.Get()))
	if err != nil {
		log.Fatalf("rpc.NewClient failed: %v", err)
		return nil
	}
	return &Helloworld{client: pb.NewGreeterClient(c.Conn())}
}

type Helloworld struct {
	client pb.GreeterClient
}

func (this *Helloworld) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := this.client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return resp, err
	}

	return resp, nil
}
