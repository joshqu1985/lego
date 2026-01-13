package redis

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

const (
	RedisNil = goredis.Nil //nolint:errname
)

type (
	Pipeliner = goredis.Pipeliner
	Z         = goredis.Z

	Store struct {
		singel  *goredis.Client
		cluster *goredis.ClusterClient
		mode    int // 0 单机 1 集群
	}

	Config struct {
		Pass      string   `json:"pass"      toml:"pass"      yaml:"pass"`
		Endpoints []string `json:"endpoints" toml:"endpoints" yaml:"endpoints"`
		Mode      int      `json:"mode"      toml:"mode"      yaml:"mode"`
		Database  int      `json:"database"  toml:"database"  yaml:"database"`
		TLS       bool     `json:"tls"       toml:"tls"       yaml:"tls"`
	}
)

// New 初始化Redis.
func New(conf Config) (*Store, error) {
	if len(conf.Endpoints) == 0 {
		return nil, errors.New("endpoints is empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	store := &Store{mode: conf.Mode}

	if conf.Mode == 0 {
		store.singel, err = newSingel(ctx, conf)
	} else {
		store.cluster, err = newCluster(ctx, conf)
	}

	return store, err
}

func (s *Store) client() goredis.Cmdable {
	if s.mode == 0 {
		return s.singel
	} else {
		return s.cluster
	}
}

func newSingel(ctx context.Context, conf Config) (*goredis.Client, error) {
	var tlsConfig *tls.Config
	if conf.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec
		}
	}

	rdb := goredis.NewClient(&goredis.Options{
		Addr:         conf.Endpoints[0],
		Password:     conf.Pass,
		DB:           conf.Database,
		MaxRetries:   3,
		MinIdleConns: 8,
		TLSConfig:    tlsConfig,
	})

	return rdb, rdb.Ping(ctx).Err()
}

func newCluster(ctx context.Context, conf Config) (*goredis.ClusterClient, error) {
	var tlsConfig *tls.Config
	if conf.TLS {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec
		}
	}

	rdb := goredis.NewClusterClient(&goredis.ClusterOptions{
		Addrs:        conf.Endpoints,
		Password:     conf.Pass,
		MaxRetries:   3,
		MinIdleConns: 8,
		TLSConfig:    tlsConfig,
	})

	return rdb, rdb.Ping(ctx).Err()
}
