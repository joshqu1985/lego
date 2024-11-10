package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// -- GenericCmdable --

func (this *Store) Exists(ctx context.Context, key string) (int64, error) {
	return this.client().Exists(ctx, key).Result() // 1: 存在 0: 不存在
}

func (this *Store) Del(ctx context.Context, key string) (int64, error) {
	return this.client().Del(ctx, key).Result() // 被删除key的数量
}

func (this *Store) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return this.client().Expire(ctx, key, expiration).Err()
}

func (this *Store) TTL(ctx context.Context, key string) (time.Duration, error) {
	return this.client().TTL(ctx, key).Result() // -2: key不存在 -1: 没设置剩余时间 n: 剩余秒
}

// -- StringCmdable --

func (this *Store) Get(ctx context.Context, key string, val any) error {
	data, err := this.client().Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return Scan(data, val)
}

func (this *Store) Set(ctx context.Context, key string, val any, args ...int64) error {
	data, err := Format(val)
	if err != nil {
		return err
	}

	ttl := time.Duration(0) * time.Second
	if len(args) > 0 {
		ttl = time.Duration(args[0]) * time.Second
	}
	return this.client().Set(ctx, key, data, ttl).Err()
}

func (this *Store) SetNX(ctx context.Context, key string, val any, args ...int64) error {
	data, err := Format(val)
	if err != nil {
		return err
	}

	ttl := time.Duration(0) * time.Second
	if len(args) > 0 {
		ttl = time.Duration(args[0]) * time.Second
	}
	return this.client().SetNX(ctx, key, data, ttl).Err()
}

func (this *Store) IncrBy(ctx context.Context, key string, val int64) (int64, error) {
	return this.client().IncrBy(ctx, key, val).Result() // val: 执行后key的值
}

func (this *Store) DecrBy(ctx context.Context, key string, val int64) (int64, error) {
	return this.client().DecrBy(ctx, key, val).Result() // val: 执行后key的值
}

// -- HashCmdable --

func (this *Store) HExists(ctx context.Context, key, field string) (bool, error) {
	return this.client().HExists(ctx, key, field).Result() // true: 存在 false: 不存在
}

func (this *Store) HGet(ctx context.Context, key, field string, val any) error {
	data, err := this.client().HGet(ctx, key, field).Bytes()
	if err != nil {
		return err
	}
	return Scan(data, val)
}

func (this *Store) HSet(ctx context.Context, key, field string, val any) error {
	data, err := Format(val)
	if err != nil {
		return err
	}
	return this.client().HSet(ctx, key, field, data).Err()
}

func (this *Store) HSetNX(ctx context.Context, key, field string, val any) error {
	data, err := Format(val)
	if err != nil {
		return err
	}
	return this.client().HSetNX(ctx, key, field, data).Err()
}

func (this *Store) HMSet(ctx context.Context, key string, vals map[string]any) error {
	values := map[string][]byte{}
	for field, val := range vals {
		data, err := Format(val)
		if err != nil {
			return err
		}
		values[field] = data
	}
	return this.client().HMSet(ctx, key, values).Err()
}

func (this *Store) HMGet(ctx context.Context, key string, fields []string, vals any) error {
	results, err := this.client().HMGet(ctx, key, fields...).Result()
	if err != nil {
		return err
	}
	values := make([]string, 0, len(fields))
	for i := 0; i < len(results); i++ {
		v, ok := results[i].(string)
		if !ok {
			return fmt.Errorf("invalid value type: %v", results[i])
		}
		values = append(values, v)
	}
	return ScanSlice(values, vals)
}

func (this *Store) HVals(ctx context.Context, key string, vals any) error {
	results, err := this.client().HVals(ctx, key).Result()
	if err != nil {
		return err
	}
	return ScanSlice(results, vals)
}

func (this *Store) HDel(ctx context.Context, key string, field ...string) (int64, error) {
	return this.client().HDel(ctx, key, field...).Result() // val: 被成功删除key的数量
}

func (this *Store) HIncrBy(ctx context.Context, key, field string, val int64) (int64, error) {
	return this.client().HIncrBy(ctx, key, field, val).Result() // val: 执行后字段的值
}

func (this *Store) HLen(ctx context.Context, key string) (int64, error) {
	return this.client().HLen(ctx, key).Result() // val: 字段的数量
}

// -- SortedSetCmdable --

type Pair struct {
	Member string
	Score  float64
}

func (this *Store) ZAdd(ctx context.Context, key string, pairs ...Pair) (int64, error) {
	members := make([]goredis.Z, 0, len(pairs))
	for _, pair := range pairs {
		members = append(members, goredis.Z{Score: pair.Score, Member: pair.Member})
	}

	return this.client().ZAdd(ctx, key, members...).Result() // val: 成功添加的成员数量
}

func (this *Store) ZCard(ctx context.Context, key string) (int64, error) {
	return this.client().ZCard(ctx, key).Result() // val: 返回集合个数
}

func (this *Store) ZCount(ctx context.Context, key string, min, max float64) (int64, error) {
	return this.client().ZCount(ctx, key, strconv.FormatFloat(min, 'f', -1, 64),
		strconv.FormatFloat(max, 'f', -1, 64)).Result() // val: 区间内的成员数量
}

func (this *Store) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Pair, error) {
	results, err := this.client().ZRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	return this.toPairs(results), nil
}

func (this *Store) ZRangeByScoreWithScores(ctx context.Context, key string, min, max float64) ([]Pair, error) {
	results, err := this.client().ZRangeByScoreWithScores(ctx, key, &goredis.ZRangeBy{
		Min: strconv.FormatFloat(min, 'f', -1, 64),
		Max: strconv.FormatFloat(max, 'f', -1, 64),
	}).Result()
	if err != nil {
		return nil, err
	}

	return this.toPairs(results), nil
}

func (this *Store) toPairs(results []goredis.Z) []Pair {
	pairs := make([]Pair, 0, len(results))
	for _, result := range results {
		member, ok := (result.Member).(string)
		if !ok {
			continue
		}
		pairs = append(pairs, Pair{Member: member, Score: result.Score})
	}
	return pairs
}

func (this *Store) ZScore(ctx context.Context, key, member string) (float64, error) {
	return this.client().ZScore(ctx, key, member).Result() // val: 成员的分数
}

func (this *Store) ZIncrBy(ctx context.Context, key, member string, incr float64) (float64, error) {
	return this.client().ZIncrBy(ctx, key, incr, member).Result() // val: 执行后的分数
}

func (this *Store) ZRem(ctx context.Context, key string, members ...any) (int64, error) {
	return this.client().ZRem(ctx, key, members...).Result() // val: 成功删除的成员数量
}

// -- ListCmdable --

func (this *Store) LLen(ctx context.Context, key string) (int64, error) {
	return this.client().LLen(ctx, key).Result() // val: 列表的长度
}

func (this *Store) LPop(ctx context.Context, key string, val any) error {
	data, err := this.client().LPop(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return Scan(data, val)
}

func (this *Store) LPush(ctx context.Context, key string, vals ...any) (int64, error) {
	values := []any{}
	for _, val := range vals {
		data, err := Format(val)
		if err != nil {
			return 0, err
		}
		values = append(values, data)
	}
	return this.client().LPush(ctx, key, values...).Result() // val: 执行后列表长度
}

func (this *Store) RPop(ctx context.Context, key string, val any) error {
	data, err := this.client().RPop(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return Scan(data, val)
}

func (this *Store) Rpush(ctx context.Context, key string, vals ...any) (int64, error) {
	values := []any{}
	for _, val := range vals {
		data, err := Format(val)
		if err != nil {
			return 0, err
		}
		values = append(values, data)
	}
	return this.client().RPush(ctx, key, values...).Result() // val: 执行后列表长度
}

// -- Cmdable --

func (this *Store) Pipelined(ctx context.Context, fn func(Pipeliner) error) error {
	_, err := this.client().Pipelined(ctx, fn)
	return err
}
