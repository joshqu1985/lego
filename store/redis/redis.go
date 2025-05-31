package redis

import (
	"context"
	"crypto/tls"

	goredis "github.com/redis/go-redis/v9"
)

const (
	RedisNil = goredis.Nil
)

type (
	Pipeliner = goredis.Pipeliner
	Z         = goredis.Z
)

type Store struct {
	singel  *goredis.Client
	cluster *goredis.ClusterClient
	mode    int // 0 单机 1 集群
}

// New 初始化Redis
func New(conf Config) *Store {
	if len(conf.Endpoints) == 0 {
		panic("endpoints is empty")
	}

	var err error
	store := &Store{mode: conf.Mode}

	if conf.Mode == 0 {
		store.singel, err = newSingel(conf)
	} else {
		store.cluster, err = newCluster(conf)
	}
	if err != nil {
		panic(err)
	}
	return store
}

func (this *Store) client() goredis.Cmdable {
	if this.mode == 0 {
		return this.singel
	} else {
		return this.cluster
	}
}

type Config struct {
	Endpoints []string `toml:"endpoints" yaml:"endpoints" json:"endpoints"`
	Auth      string   `toml:"auth" yaml:"pass" json:"auth"`
	Database  int      `toml:"database" yaml:"database" json:"database"`
	TLS       bool     `toml:"tls" yaml:"tls" json:"tls"`
	Mode      int      `toml:"mode" yaml:"mode" json:"mode"` // 0 单机 1 集群
}

func newSingel(conf Config) (*goredis.Client, error) {
	var tlsConfig *tls.Config
	if conf.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	rdb := goredis.NewClient(&goredis.Options{
		Addr:         conf.Endpoints[0],
		Password:     conf.Auth,
		DB:           conf.Database,
		MaxRetries:   3,
		MinIdleConns: 8,
		TLSConfig:    tlsConfig,
	})
	return rdb, rdb.Ping(context.Background()).Err()
}

func newCluster(conf Config) (*goredis.ClusterClient, error) {
	var tlsConfig *tls.Config
	if conf.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	rdb := goredis.NewClusterClient(&goredis.ClusterOptions{
		Addrs:        conf.Endpoints,
		Password:     conf.Auth,
		MaxRetries:   3,
		MinIdleConns: 8,
		TLSConfig:    tlsConfig,
	})
	return rdb, rdb.Ping(context.Background()).Err()
}
