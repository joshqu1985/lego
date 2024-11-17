package main

import (
	"context"
	"log"

	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
)

type HelloService struct {
	pb.UnimplementedGreeterServer
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

// SayHello implements helloworld.GreeterServer
func (this *HelloService) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
