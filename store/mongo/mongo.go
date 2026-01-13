package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"          //nolint:gomodguard
	"go.mongodb.org/mongo-driver/v2/mongo/options"  //nolint:gomodguard
	"go.mongodb.org/mongo-driver/v2/mongo/readpref" //nolint:gomodguard

	"github.com/joshqu1985/lego/metrics"
)

type (
	// DB 避免业务代码直接引用mongo sdk.
	DB = mongo.Database

	Store struct {
		client *mongo.Client
		db     *mongo.Database
	}

	// Config mongo配置.
	Config struct {
		Endpoint    string `json:"endpoint"      toml:"endpoint"      yaml:"endpoint"`
		User        string `json:"user"          toml:"user"          yaml:"user"`
		Pass        string `json:"pass"          toml:"pass"          yaml:"pass"`
		Opts        string `json:"opts"          toml:"opts"          yaml:"opts"`
		Database    string `json:"database"      toml:"database"      yaml:"database"`
		MaxPoolSize uint64 `json:"max_pool_size" toml:"max_pool_size" yaml:"max_pool_size"`
		MinPoolSize uint64 `json:"min_pool_size" toml:"min_pool_size" yaml:"min_pool_size"`
		MaxLife     int64  `json:"max_life"      toml:"max_life"      yaml:"max_life"`
	}
)

var (
	metricsDuration = metrics.NewHistogram(&metrics.HistogramOpt{
		Namespace: "mongo_client",
		Name:      "exec_duration",
		Help:      "mongo exec duration",
		Labels:    []string{"command"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	metricsSlowTotal = metrics.NewCounter(&metrics.CounterOpt{
		Namespace: "mongo_client",
		Name:      "slow_total",
		Help:      "mongo slow total",
		Labels:    []string{"command"},
	})
)

func (s *Store) Exec(fn func(context.Context, *DB) error, labels ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now()
	err := fn(ctx, s.db)

	// metrics
	elapse := time.Since(now).Milliseconds()
	if len(labels) != 0 {
		metricsDuration.Observe(elapse, labels...)
	}
	if len(labels) != 0 && elapse > 500 {
		metricsSlowTotal.Inc(labels...)
	}

	return err
}

func (s *Store) Transaction(fn func(context.Context, *DB) error, labels ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now()
	session, err := s.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if xerr := session.StartTransaction(); xerr != nil {
		return xerr
	}

	_, err = session.WithTransaction(ctx, func(sessCtx context.Context) (any, error) {
		return nil, fn(sessCtx, s.db)
	})

	// metrics
	elapse := time.Since(now).Milliseconds()
	if len(labels) != 0 {
		metricsDuration.Observe(elapse, labels...)
	}
	if len(labels) != 0 && elapse > 500 {
		metricsSlowTotal.Inc(labels...)
	}

	return err
}

func New(conf *Config) (*Store, error) {
	store := &Store{}

	client, err := connect(conf)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if xerr := ping(ctx, client); xerr != nil {
		return nil, xerr
	}

	store.client = client
	store.db = client.Database(conf.Database)

	return store, nil
}

func connect(conf *Config) (*mongo.Client, error) {
	dsn := "mongodb://"
	if conf.User != "" && conf.Pass != "" {
		dsn = dsn + conf.User + ":" + conf.Pass + "@"
	}

	dsn += conf.Endpoint
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

func ping(ctx context.Context, client *mongo.Client) error {
	return client.Ping(ctx, readpref.Primary())
}
