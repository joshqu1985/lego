package mysql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joshqu1985/lego/metrics"
)

type Store struct {
	db *DB
}

func (this *Store) Exec(fn func(*DB) error, labels ...string) error {
	now := time.Now()
	err := fn(this.db)

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

func (this *Store) Transaction(fn func(*DB) error, labels ...string) error {
	now := time.Now()
	err := this.db.Transaction(fn)

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

func (this *Store) Begin() *DB {
	return this.db.Begin()
}

func (this *Store) Rollback() *DB {
	return this.db.Rollback()
}

func (this *Store) Commit() *DB {
	return this.db.Commit()
}

// Config mysql配置
type Config struct {
	Endpoint     string `toml:"endpoint" yaml:"endpoint" json:"endpoint"` // 10.X.X.X:3306
	Auth         string `toml:"auth" yaml:"auth" json:"auth"`
	Opts         string `toml:"opts" yaml:"opts" json:"opts"`
	Database     string `toml:"database" yaml:"database" json:"database"`
	MaxOpenConns int    `toml:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxLife      int    `toml:"max_life" yaml:"max_life" json:"max_life"` // ms
}

// New 初始化mysql连接池
func New(conf Config) *Store {
	store := &Store{}

	var err error
	store.db, err = connect(conf)
	if err != nil {
		panic(err)
	}

	log.Printf("mysql init host:%s finish\n", conf.Endpoint)
	return store
}

func connect(conf Config) (*DB, error) {
	dsn := fmt.Sprintf("%s@tcp(%s)/%s", conf.Auth, conf.Endpoint, conf.Database)
	if conf.Opts != "" {
		dsn = dsn + "?" + conf.Opts
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(conf.MaxIdleConns)
	conn.SetMaxOpenConns(conf.MaxOpenConns)
	if conf.MaxLife != 0 {
		conn.SetConnMaxLifetime(time.Duration(conf.MaxLife) * time.Millisecond)
	}
	return db, nil
}

var (
	metricsDuration = metrics.NewHistogram(metrics.HistogramOpt{
		Namespace: "mysql_client",
		Name:      "duration_ms",
		Labels:    []string{"command"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	metricsSlowCount = metrics.NewCounter(metrics.CounterOpt{
		Namespace: "mysql_client",
		Name:      "slow_count",
		Labels:    []string{"command"},
	})
)
