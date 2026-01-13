package main

import (
	"github.com/joshqu1985/lego/configor"
	"github.com/joshqu1985/lego/logs"
	"github.com/joshqu1985/lego/store/mysql"
	"github.com/joshqu1985/lego/store/redis"
	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rest"

	"{{$.ServerName}}/httpserver"
	"{{$.ServerName}}/internal/config"
	"{{$.ServerName}}/internal/repo"
	"{{$.ServerName}}/internal/service"
)

var (
  n naming.Naming
)

func init() {
	if err := configor.New("./config.toml",
		configor.WithEtcd(),
		configor.WithToml(),
	).Load(&config.C); err != nil {
		panic(err)
	}

	if n, err = naming.Init(config.C.Naming); err != nil {
		panic(err)
	}

	logs.Init(logs.WithDevelopment(), logs.WithWriterConsole())
}

func main() {
	svc := service.New(
		repo.New(mysql.New(config.C.Mysql), redis.New(config.C.Redis)),
	)

	s, err := rest.NewServer(config.C.Name, config.C.Addr,
		rest.WithNaming(n),
		rest.WithMetrics(),
		rest.WithRouters(httpserver.New(svc).Routers))
	if err != nil {
		logs.Fatalf("new rest server err:%v", err)
		return
	}

	_ = s.Start()
}
