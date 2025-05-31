package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"github.com/joshqu1985/lego/metrics"
)

type (
	// DB 避免业务代码直接引用mongo sdk
	DB = mongo.Database
)

type Store struct {
	client *mongo.Client
	db     *mongo.Database
}

func (this *Store) Exec(fn func(context.Context, *DB) error, labels ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now()
	err := fn(ctx, this.db)

	// metrics
	elapse := time.Since(now).Milliseconds()
	if len(labels) != 0 {
		metricsDuration.Observe(elapse, labels...)
	}
	if len(labels) != 0 && elapse > 500 {
		metricsSlowCount.Inc(labels...)
	}
	return err
}

func (this *Store) Transaction(fn func(context.Context, *DB) error, labels ...string) error {
	now := time.Now()
	session, err := this.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	if err := session.StartTransaction(); err != nil {
		return err
	}

	_, err = session.WithTransaction(context.Background(), func(sessCtx context.Context) (any, error) {
		return nil, fn(sessCtx, this.db)
	})

	// metrics
	elapse := time.Since(now).Milliseconds()
	if len(labels) != 0 {
		metricsDuration.Observe(elapse, labels...)
	}
	if len(labels) != 0 && elapse > 500 {
		metricsSlowCount.Inc(labels...)
	}
	return err
}

// Config mongo配置
type Config struct {
	Endpoint    string `toml:"endpoint" yaml:"endpoint" json:"endpoint"` // 10.X.X.X:27017,10.X.X.X:27017/admin
	Auth        string `toml:"auth" yaml:"auth" json:"auth"`             // user:pass
	Opts        string `toml:"opts" yaml:"opts" json:"opts"`             // replicaSet=cmgo-*******&authSource=admin
	Database    string `toml:"database" yaml:"database" json:"database"` //
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
	if err := ping(client); err != nil {
		panic(err)
	}

	store.client = client
	store.db = client.Database(conf.Database)

	log.Printf("mongo init host:%s finish\n", conf.Endpoint)
	return store
}

func connect(conf Config) (*mongo.Client, error) {
	dsn := "mongodb://"
	if conf.Auth != "" {
		dsn = dsn + conf.Auth + "@"
	}

	dsn = dsn + conf.Endpoint
	if conf.Opts != "" {
		dsn = dsn + "?" + conf.Opts
	}

	clientOpts := options.Client().ApplyURI(dsn)

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

func ping(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return client.Ping(ctx, readpref.Primary())
}

var (
	metricsDuration = metrics.NewHistogram(metrics.HistogramOpt{
		Namespace: "mongo_client",
		Name:      "duration_ms",
		Labels:    []string{"command"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	metricsSlowCount = metrics.NewCounter(metrics.CounterOpt{
		Namespace: "mongo_client",
		Name:      "slow_count",
		Labels:    []string{"command"},
	})
)
