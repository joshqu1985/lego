package main

import (
	"log"

	"google.golang.org/grpc"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/route_guide/routeguide"
)

func init() {
	_, _ = naming.Init(naming.Config{
		// Source:    "etcd",
		// Endpoints: []string{"127.0.0.1:9379"},
		Source:    "nacos",
		Endpoints: []string{"127.0.0.1:8848"},
		Cluster:   "f5fc10a5-0f22-4188-ac9c-1e34023f6556",
	})
}

func main() {
	config := rpc.Config{
		Name: "route_guide",
		Addr: "127.0.0.1:50053",
	}

	handlers := func(s *grpc.Server) {
		pb.RegisterRouteGuideServer(s, NewRouteGuide())
	}

	s, err := rpc.NewServer(config, rpc.WithNaming(naming.Get()),
		rpc.WithHandlers(handlers))
	if err != nil {
		log.Fatalf("new rpc server err:%v", err)
		return
	}

	_ = s.Start()
}
