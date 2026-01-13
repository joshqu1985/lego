package redis

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	rdb  *Store
	once sync.Once
)

func InitTestRedis() {
	once.Do(func() {
		rdb, _ = New(Config{
			Endpoints: []string{"127.0.0.1:6379"},
		})
	})
}

type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Test_StringCmdable(t *testing.T) {
	InitTestRedis()
	ctx := context.Background()

	// string Get Set
	strVal1 := "test_val"
	var strVal2 string
	err := rdb.Set(ctx, "test_string", strVal1)
	assert.Nil(t, err)
	err = rdb.Get(ctx, "test_string", &strVal2)
	assert.Nil(t, err)
	assert.Equal(t, strVal1, strVal2)

	// struct Get Set
	st1 := TestStruct{
		Name: "test_name",
		Age:  18,
	}
	var st2 TestStruct
	err = rdb.Set(ctx, "test_struct", &st1)
	assert.Nil(t, err)
	err = rdb.Get(ctx, "test_struct", &st2)
	assert.Nil(t, err)
	assert.Equal(t, st1, st2)

	// time Get Set
	t1 := time.Now()
	var t2 time.Time
	err = rdb.Set(ctx, "test_time", t1)
	assert.Nil(t, err)
	err = rdb.Get(ctx, "test_time", &t2)
	assert.Nil(t, err)
	assert.Equal(t, t1.UnixNano(), t2.UnixNano())

	// int IncrBy DecrBy
	val, err := rdb.IncrBy(ctx, "test_int", 10)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), val)

	var data int
	err = rdb.Get(ctx, "test_int", &data)
	assert.Nil(t, err)
	assert.Equal(t, 10, data)

	val, err = rdb.DecrBy(ctx, "test_int", 10)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), val)
}

func Test_HashCmdable(t *testing.T) {
	InitTestRedis()
	ctx := context.Background()

	// string HSet HGet
	hashVal1 := "hash_val1"
	err := rdb.HSet(ctx, "test_hash_string", "field1", hashVal1)
	assert.Nil(t, err)

	var hashData1 string
	err = rdb.HGet(ctx, "test_hash_string", "field1", &hashData1)
	assert.Nil(t, err)
	assert.Equal(t, hashVal1, hashData1)

	// struct HSet HGet
	hashVal2 := TestStruct{
		Name: "test_name",
		Age:  18,
	}
	err = rdb.HSet(ctx, "test_hash_struct", "field2", hashVal2)
	assert.Nil(t, err)

	var hashData2 TestStruct
	err = rdb.HGet(ctx, "test_hash_struct", "field2", &hashData2)
	assert.Nil(t, err)
	assert.Equal(t, hashVal2, hashData2)
}

func Test_Pipelined(t *testing.T) {
	InitTestRedis()
	ctx := context.Background()

	err := rdb.Pipelined(ctx,
		func(pipe Pipeliner) error {
			pipe.Incr(ctx, "test_pipeline_incr")
			pipe.ZAdd(ctx, "test_pipeline_zadd", Z{Member: "m1", Score: 12})
			return nil
		},
	)
	assert.Nil(t, err)
}
