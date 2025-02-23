package main

import (
	"github.com/joshqu1985/lego/configor"
	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/store/mysql"
	"github.com/joshqu1985/lego/store/redis"
	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rpc"

	"{{$.ServerName}}/internal/config"
	"{{$.ServerName}}/internal/repo"
	"{{$.ServerName}}/internal/service"
	"{{$.ServerName}}/rpcserver"
)

func init() {
	if err := configor.New("./config.toml",
		configor.WithEtcd(),
		configor.WithToml(),
	).Load(&config.C); err != nil {
		panic(err)
	}

	if _, err := naming.Init(config.C.Naming); err != nil {
		panic(err)
	}

	logs.Init(logs.WithDevelopment(), logs.WithWriterConsole())
}

func main() {
	svc := service.New(
		repo.New(mysql.New(config.C.Mysql), redis.New(config.C.Redis)),
	)

	s, err := rpc.NewServer(config.C.Name, config.C.Addr,
		rpc.WithNaming(naming.Get()),
		rpc.WithMetrics(),
		rpc.WithRouters(rpcserver.New(svc).Routers))
	if err != nil {
		logs.Fatalf("new rest server err:%v", err)
	}

	_ = s.Start()
}
