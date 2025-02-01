package main

import (
	"google.golang.org/grpc"

	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
)

func RouterRegister(s *grpc.Server) {
	pb.RegisterGreeterServer(s, NewHelloService())
}
