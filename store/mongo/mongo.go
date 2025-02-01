package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type (
	// DB 避免业务代码直接引用mongo sdk
	DB = mongo.Database
)

type Store struct {
	client *mongo.Client
	db     *mongo.Database
}

func (this *Store) Exec(ctx context.Context, fn func(*DB) error) error {
	err := fn(this.db)

	return err
}

// Config mongo配置
type Config struct {
	Endpoint    string `toml:"endpoint" yaml:"endpoint" json:"endpoint"`
	Username    string `toml:"username" yaml:"username" json:"username"`
	Password    string `toml:"password" yaml:"password" json:"password"`
	Database    string `toml:"database" yaml:"database" json:"database"`
	MaxPoolSize uint64 `toml:"max_pool_size" yaml:"max_pool_size" json:"max_pool_size"`
	MinPoolSize uint64 `toml:"min_pool_size" yaml:"min_pool_size" json:"min_pool_size"`
	MaxLife     uint64 `toml:"max_life" yaml:"max_life" json:"max_life"` // ms
}

func New(conf Config) *Store {
	store := &Store{}

	client, err := connect(conf)
	if err != nil {
		panic(err)
	}
	store.client = client
	store.db = client.Database(conf.Database)

	log.Printf("mongo init host:%s finish\n", conf.Endpoint)
	return store
}

func connect(conf Config) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(conf.Endpoint)
	credential := options.Credential{
		Username: conf.Username,
		Password: conf.Password,
	}
	if conf.Username != "" && conf.Password != "" {
		clientOpts.SetAuth(credential)
	}
	if conf.MaxPoolSize != 0 {
		clientOpts.SetMaxPoolSize(conf.MaxPoolSize)
	}
	if conf.MinPoolSize != 0 {
		clientOpts.SetMinPoolSize(conf.MinPoolSize)
	}
	if conf.MaxLife != 0 {
		clientOpts.SetMaxConnIdleTime(time.Duration(conf.MaxLife) * time.Millisecond)
	}
	return mongo.Connect(clientOpts)
}
