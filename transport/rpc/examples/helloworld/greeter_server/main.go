package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	// "net"

	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
	"google.golang.org/grpc"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	n := naming.New(naming.Config{
		Endpoints: []string{"http://127.0.0.1:9379"},
		Cluster:   "default",
	})

	register := func(s *grpc.Server) {
		pb.RegisterGreeterServer(s, &server{})
	}

	s, err := rpc.NewServer(rpc.ServerConfig{
		Name: "helloword",
		Addr: "127.0.0.1:50051",
	}, rpc.WithNaming(n), rpc.WithRegisterFunc(register))
	if err != nil {
		fmt.Println("--", err)
		return
	}
	s.Start()
}

/*
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
*/
