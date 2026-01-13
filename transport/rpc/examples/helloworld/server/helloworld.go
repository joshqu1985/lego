package main

import (
	"context"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/joshqu1985/lego/logs"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
)

type HelloService struct {
	pb.UnimplementedGreeterServer
}

var count int64

func NewHelloService() *HelloService {
	return &HelloService{}
}

// SayHello implements helloworld.GreeterServer.
func (hs *HelloService) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	atomic.AddInt64(&count, 1)
	if atomic.LoadInt64(&count)%10000 == 0 {
		logs.Info("-------->", atomic.LoadInt64(&count)/10000, "万")
	}
	if atomic.LoadInt64(&count) > 500000 && atomic.LoadInt64(&count) < 2000000 {
		return nil, status.New(codes.Internal, "server busy").Err()
	}

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
