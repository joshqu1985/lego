package orm

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"    // nolint:gomodguard
	"gorm.io/driver/postgres" // nolint:gomodguard
	"gorm.io/gorm"            // nolint:gomodguard

	"github.com/joshqu1985/lego/metrics"
)

const (
	DB_MYSQL   = "mysql"
	DB_POSTGRE = "postgre"
)

type (
	Store struct {
		durationHandle  metrics.HistogramVec
		slowCountHandle metrics.CounterVec
		db              *DB
		source          string
	}

	// Config mysql配置.
	Config struct {
		Source       string `json:"source"         toml:"source"         yaml:"source"`
		Host         string `json:"host"           toml:"host"           yaml:"host"`
		User         string `json:"user"           toml:"user"           yaml:"user"`
		Pass         string `json:"pass"           toml:"pass"           yaml:"pass"`
		Opts         string `json:"opts"           toml:"opts"           yaml:"opts"`
		Database     string `json:"database"       toml:"database"       yaml:"database"`
		Port         int    `json:"port"           toml:"port"           yaml:"port"`
		MaxOpenConns int    `json:"max_open_conns" toml:"max_open_conns" yaml:"max_open_conns"`
		MaxIdleConns int    `json:"max_idle_conns" toml:"max_idle_conns" yaml:"max_idle_conns"`
		MaxLife      int    `json:"max_life"       toml:"max_life"       yaml:"max_life"`
	}
)

var (
	mysqlDuration = metrics.NewHistogram(&metrics.HistogramOpt{
		Namespace: DB_MYSQL,
		Name:      "exec_duration",
		Help:      "mysql exec duration",
		Labels:    []string{"command"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	mysqlSlowTotal = metrics.NewCounter(&metrics.CounterOpt{
		Namespace: DB_MYSQL,
		Name:      "slow_total",
		Help:      "mysql slow total",
		Labels:    []string{"command"},
	})

	postgreDuration = metrics.NewHistogram(&metrics.HistogramOpt{
		Namespace: DB_POSTGRE,
		Name:      "exec_duration",
		Help:      "postgre exec duration",
		Labels:    []string{"command"},
		Buckets:   []float64{3, 5, 10, 50, 100, 250, 500, 1000, 2000, 5000},
	})

	postgreSlowTotal = metrics.NewCounter(&metrics.CounterOpt{
		Namespace: DB_POSTGRE,
		Name:      "slow_total",
		Help:      "postgre slow total",
		Labels:    []string{"command"},
	})
)

func (s *Store) Exec(fn func(*DB) error, labels ...string) error {
	now := time.Now()
	err := fn(s.db)
	elapse := time.Since(now).Milliseconds()

	// metrics
	if s.durationHandle != nil && len(labels) != 0 {
		s.durationHandle.Observe(elapse, labels...)
	}
	if s.slowCountHandle != nil && len(labels) != 0 && elapse > 500 {
		s.slowCountHandle.Inc(labels...)
	}

	return err
}

func (s *Store) Transaction(fn func(*DB) error, labels ...string) error {
	now := time.Now()
	err := s.db.Transaction(fn)
	elapse := time.Since(now).Milliseconds()

	// metrics
	if s.durationHandle != nil && len(labels) != 0 {
		s.durationHandle.Observe(elapse, labels...)
	}
	if s.slowCountHandle != nil && len(labels) != 0 && elapse > 500 {
		s.slowCountHandle.Inc(labels...)
	}

	return err
}

func (s *Store) Begin() *DB {
	return s.db.Begin()
}

func (s *Store) Rollback() *DB {
	return s.db.Rollback()
}

func (s *Store) Commit() *DB {
	return s.db.Commit()
}

// New 初始化orm连接池.
func New(conf *Config) (*Store, error) {
	store := &Store{source: conf.Source}

	var err error
	switch conf.Source {
	case DB_MYSQL:
		store.durationHandle, store.slowCountHandle = mysqlDuration, mysqlSlowTotal
		store.db, err = connect(conf)
	case DB_POSTGRE:
		store.durationHandle, store.slowCountHandle = postgreDuration, postgreSlowTotal
		store.db, err = connect(conf)
	default:
		store.db, err = nil, errors.New("source not support")
	}

	return store, err
}

func connect(conf *Config) (*DB, error) {
	var dial gorm.Dialector
	switch conf.Source {
	case DB_MYSQL:
		dial = dialectorMysql(conf)
	case DB_POSTGRE:
		dial = dialectorPostgre(conf)
	}

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		return nil, err
	}
	if xerr := conn.Ping(); xerr != nil {
		return nil, xerr
	}

	conn.SetMaxIdleConns(conf.MaxIdleConns)
	conn.SetMaxOpenConns(conf.MaxOpenConns)
	if conf.MaxLife != 0 {
		conn.SetConnMaxLifetime(time.Duration(conf.MaxLife) * time.Millisecond)
	}

	return db, nil
}

func dialectorMysql(conf *Config) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	if conf.Opts != "" {
		return mysql.Open(dsn + "?" + conf.Opts) // charset=utf8mb4&parseTime=True
	}

	return mysql.Open(dsn)
}

func dialectorPostgre(conf *Config) gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Database,
		conf.Port,
	)

	if conf.Opts != "" {
		return postgres.Open(dsn + " " + conf.Opts) // sslmode=disable TimeZone=Asia/Shanghai
	}

	return postgres.Open(dsn)
}
