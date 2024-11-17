package main

import (
	"context"
	"log"
	"time"

	"github.com/joshqu1985/lego/transport/naming"
	pb "github.com/joshqu1985/lego/transport/rpc/examples/helloworld/helloworld"
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
	client := NewHelloworldClient(naming.Get())

	for i := 0; i < 20; i++ {
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
		log.Printf("Greeting: resp:%s err:%v", resp.GetMessage(), err)

		time.Sleep(time.Millisecond * 500)
	}
}
