package main

import (
	"log"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rest"
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

	name := "helloworld"
	addr := "127.0.0.1:50051"

	s, err := rest.NewServer(name, addr, rest.WithRouters(RouterRegister),
		rest.WithNaming(n), rest.WithMetrics())
	if err != nil {
		log.Fatalf("new rest server err:%v", err)

		return
	}

	_ = s.Start()
}
