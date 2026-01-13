package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Pair struct {
	Member string
	Score  float64
}

// -- GenericCmdable --

func (s *Store) Exists(ctx context.Context, key string) (int64, error) {
	return s.client().Exists(ctx, key).Result() // 1: 存在 0: 不存在
}

func (s *Store) Del(ctx context.Context, key string) (int64, error) {
	return s.client().Del(ctx, key).Result() // 被删除key的数量
}

func (s *Store) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return s.client().Expire(ctx, key, expiration).Err()
}

func (s *Store) TTL(ctx context.Context, key string) (time.Duration, error) {
	return s.client().TTL(ctx, key).Result() // -2: key不存在 -1: 没设置剩余时间 n: 剩余秒
}

// -- StringCmdable --

func (s *Store) Get(ctx context.Context, key string, val any) error {
	data, err := s.client().Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return Scan(data, val)
}

func (s *Store) Set(ctx context.Context, key string, val any, args ...int64) error {
	data, err := Format(val)
	if err != nil {
		return err
	}

	ttl := time.Duration(0) * time.Second
	if len(args) > 0 {
		ttl = time.Duration(args[0]) * time.Second
	}

	return s.client().Set(ctx, key, data, ttl).Err()
}

func (s *Store) SetNX(ctx context.Context, key string, val any, args ...int64) error {
	data, err := Format(val)
	if err != nil {
		return err
	}

	ttl := time.Duration(0) * time.Second
	if len(args) > 0 {
		ttl = time.Duration(args[0]) * time.Second
	}

	return s.client().SetNX(ctx, key, data, ttl).Err()
}

func (s *Store) IncrBy(ctx context.Context, key string, val int64) (int64, error) {
	return s.client().IncrBy(ctx, key, val).Result() // val: 执行后key的值
}

func (s *Store) DecrBy(ctx context.Context, key string, val int64) (int64, error) {
	return s.client().DecrBy(ctx, key, val).Result() // val: 执行后key的值
}

// -- HashCmdable --

func (s *Store) HExists(ctx context.Context, key, field string) (bool, error) {
	return s.client().HExists(ctx, key, field).Result() // true: 存在 false: 不存在
}

func (s *Store) HGet(ctx context.Context, key, field string, val any) error {
	data, err := s.client().HGet(ctx, key, field).Bytes()
	if err != nil {
		return err
	}

	return Scan(data, val)
}

func (s *Store) HSet(ctx context.Context, key, field string, val any) error {
	data, err := Format(val)
	if err != nil {
		return err
	}

	return s.client().HSet(ctx, key, field, data).Err()
}

func (s *Store) HSetNX(ctx context.Context, key, field string, val any) error {
	data, err := Format(val)
	if err != nil {
		return err
	}

	return s.client().HSetNX(ctx, key, field, data).Err()
}

func (s *Store) HMSet(ctx context.Context, key string, vals map[string]any) error {
	values := make(map[string][]byte)
	for field, val := range vals {
		data, err := Format(val)
		if err != nil {
			return err
		}
		values[field] = data
	}

	return s.client().HMSet(ctx, key, values).Err()
}

func (s *Store) HMGet(ctx context.Context, key string, fields []string, vals any) error {
	results, err := s.client().HMGet(ctx, key, fields...).Result()
	if err != nil {
		return err
	}
	values := make([]string, 0, len(fields))
	for i := range results {
		v, ok := results[i].(string)
		if !ok {
			return fmt.Errorf("invalid value type: %v", results[i])
		}
		values = append(values, v)
	}

	return ScanSlice(values, vals)
}

func (s *Store) HVals(ctx context.Context, key string, vals any) error {
	results, err := s.client().HVals(ctx, key).Result()
	if err != nil {
		return err
	}

	return ScanSlice(results, vals)
}

func (s *Store) HDel(ctx context.Context, key string, field ...string) (int64, error) {
	return s.client().HDel(ctx, key, field...).Result() // val: 被成功删除key的数量
}

func (s *Store) HIncrBy(ctx context.Context, key, field string, val int64) (int64, error) {
	return s.client().HIncrBy(ctx, key, field, val).Result() // val: 执行后字段的值
}

func (s *Store) HLen(ctx context.Context, key string) (int64, error) {
	return s.client().HLen(ctx, key).Result() // val: 字段的数量
}

// -- SortedSetCmdable --

func (s *Store) ZAdd(ctx context.Context, key string, pairs ...Pair) (int64, error) {
	members := make([]goredis.Z, 0, len(pairs))
	for _, pair := range pairs {
		members = append(members, goredis.Z{Score: pair.Score, Member: pair.Member})
	}

	return s.client().ZAdd(ctx, key, members...).Result() // val: 成功添加的成员数量
}

func (s *Store) ZCard(ctx context.Context, key string) (int64, error) {
	return s.client().ZCard(ctx, key).Result() // val: 返回集合个数
}

func (s *Store) ZCount(ctx context.Context, key string, minVal, maxVal float64) (int64, error) {
	return s.client().ZCount(ctx, key, strconv.FormatFloat(minVal, 'f', -1, 64),
		strconv.FormatFloat(maxVal, 'f', -1, 64)).Result() // val: 区间内的成员数量
}

func (s *Store) ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]Pair, error) {
	results, err := s.client().ZRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	return s.toPairs(results), nil
}

func (s *Store) ZRangeByScoreWithScores(ctx context.Context, key string, minVal, maxVal float64) ([]Pair, error) {
	results, err := s.client().ZRangeByScoreWithScores(ctx, key, &goredis.ZRangeBy{
		Min: strconv.FormatFloat(minVal, 'f', -1, 64),
		Max: strconv.FormatFloat(maxVal, 'f', -1, 64),
	}).Result()
	if err != nil {
		return nil, err
	}

	return s.toPairs(results), nil
}

func (s *Store) toPairs(results []goredis.Z) []Pair {
	pairs := make([]Pair, 0, len(results))
	for _, result := range results {
		member, ok := result.Member.(string)
		if !ok {
			continue
		}
		pairs = append(pairs, Pair{Member: member, Score: result.Score})
	}

	return pairs
}

func (s *Store) ZScore(ctx context.Context, key, member string) (float64, error) {
	return s.client().ZScore(ctx, key, member).Result() // val: 成员的分数
}

func (s *Store) ZIncrBy(ctx context.Context, key, member string, incr float64) (float64, error) {
	return s.client().ZIncrBy(ctx, key, incr, member).Result() // val: 执行后的分数
}

func (s *Store) ZRem(ctx context.Context, key string, members ...any) (int64, error) {
	return s.client().ZRem(ctx, key, members...).Result() // val: 成功删除的成员数量
}

// -- ListCmdable --

func (s *Store) LLen(ctx context.Context, key string) (int64, error) {
	return s.client().LLen(ctx, key).Result() // val: 列表的长度
}

func (s *Store) LPop(ctx context.Context, key string, val any) error {
	data, err := s.client().LPop(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return Scan(data, val)
}

func (s *Store) LPush(ctx context.Context, key string, vals ...any) (int64, error) {
	values := make([]any, 0)
	for _, val := range vals {
		data, err := Format(val)
		if err != nil {
			return 0, err
		}
		values = append(values, data)
	}

	return s.client().LPush(ctx, key, values...).Result() // val: 执行后列表长度
}

func (s *Store) RPop(ctx context.Context, key string, val any) error {
	data, err := s.client().RPop(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return Scan(data, val)
}

func (s *Store) Rpush(ctx context.Context, key string, vals ...any) (int64, error) {
	values := make([]any, 0)
	for _, val := range vals {
		data, err := Format(val)
		if err != nil {
			return 0, err
		}
		values = append(values, data)
	}

	return s.client().RPush(ctx, key, values...).Result() // val: 执行后列表长度
}

// -- Cmdable --

func (s *Store) Pipelined(ctx context.Context, fn func(Pipeliner) error) error {
	_, err := s.client().Pipelined(ctx, fn)

	return err
}
