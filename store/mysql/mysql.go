package mysql

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	db *DB
}

func (this *Store) Exec(ctx context.Context, fn func(*DB) error) error {
	stacks := make([]uintptr, 1)
	runtime.Callers(2, stacks)

	err := fn(this.db)

	glog.Info(runtime.FuncForPC(stacks[0]).Name())
	return err
}

func (this *Store) Transaction(ctx context.Context, fn func(*DB) error) error {
	stacks := make([]uintptr, 1)
	runtime.Callers(2, stacks)

	err := this.db.Transaction(fn)

	glog.Info(runtime.FuncForPC(stacks[0]).Name())
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
	Endpoint string `toml:"endpoint" yaml:"endpoint" json:"endpoint"`
	Auth     string `toml:"auth" yaml:"auth" json:"auth"`
	Opts     string `toml:"opts" yaml:"opts" json:"opts"`
	Database string `toml:"database" yaml:"database" json:"database"`
	MaxIdle  int    `toml:"max_idle" yaml:"max_idle" json:"max_idle"`
	MaxOpen  int    `toml:"max_open" yaml:"max_open" json:"max_open"`
	MaxLife  int    `toml:"max_life" yaml:"max_life" json:"max_life"`
}

// New 初始化mysql连接池
func New(conf Config) *Store {
	store := &Store{}

	var err error
	store.db, err = connect(conf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mysql init host:%s finish\n", conf.Endpoint)
	return store
}

func connect(conf Config) (*DB, error) {
	dsn := fmt.Sprintf("%s@tcp(%s)/%s", conf.Auth, conf.Endpoint, conf.Database)
	if len(conf.Opts) > 0 {
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

	conn.SetMaxIdleConns(conf.MaxIdle)
	conn.SetMaxOpenConns(conf.MaxOpen)
	if conf.MaxLife != 0 {
		conn.SetConnMaxLifetime(time.Duration(conf.MaxLife) * time.Millisecond)
	}

	return db, nil
}
