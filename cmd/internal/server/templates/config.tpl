package config

import (
	"github.com/joshqu1985/lego/store/mysql"
	"github.com/joshqu1985/lego/store/redis"
	"github.com/joshqu1985/lego/transport/naming"
)

var C struct {
	Name   string        `toml:"name"`
	Addr   string        `toml:"addr"`
	Naming naming.Config `toml:"naming"`
	Redis  redis.Config  `toml:"redis"`
	Mysql  mysql.Config  `toml:"mysql"`
}
