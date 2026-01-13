package main

import (
	"context"

	"github.com/joshqu1985/lego/transport/naming"
)

var n naming.Naming

func Init() {
	n, _ = naming.New(&naming.Config{
		Source:    "etcd",
		Endpoints: []string{"127.0.0.1:9379"},
		// Source:    "nacos",
		// Endpoints: []string{"127.0.0.1:8848"},
		// Cluster:   "f5fc10a5-0f22-4188-ac9c-1e34023f6556",
	})
}

func main() {
	Init()

	client := NewHelloworld(n)

	for range 20000000 {
		_, _ = client.SayHello(context.Background(), " kk")
	}
}
