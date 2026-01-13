# Redis存储模块

本模块是一个基于go-redis/v9封装的Redis客户端，提供了统一的接口来操作Redis数据库，支持单机和集群两种部署模式，并提供了便捷的数据序列化和反序列化功能。

## 核心特性

- **统一接口**：提供统一的API操作Redis，避免业务代码直接依赖具体的Redis客户端库
- **多模式支持**：同时支持Redis单机模式和集群模式
- **数据结构支持**：支持String、Hash、SortedSet、List等Redis常用数据结构
- **自动序列化**：内置丰富的数据类型序列化和反序列化功能
- **连接池管理**：自动管理连接池，优化资源利用
- **管道操作**：支持Redis管道操作，提高批量操作性能
- **TLS加密**：支持TLS加密连接，保障数据传输安全

## 接口定义

### 主要结构体

#### Store

`Store`是Redis模块的核心结构体，封装了Redis客户端和操作方法：

```go
type Store struct {
    singel  *goredis.Client       // 单机模式客户端
    cluster *goredis.ClusterClient // 集群模式客户端
    mode    int                   // 0 单机 1 集群
}
```

#### Config

`Config`结构体用于配置Redis连接参数：

```go
type Config struct {
    Endpoints []string `toml:"endpoints" yaml:"endpoints" json:"endpoints"` // Redis服务器地址列表
    Pass      string   `toml:"pass" yaml:"pass" json:"pass"`               // 密码
    Database  int      `toml:"database" yaml:"database" json:"database"`   // 数据库索引(仅单机模式)
    TLS       bool     `toml:"tls" yaml:"tls" json:"tls"`                 // 是否启用TLS加密
    Mode      int      `toml:"mode" yaml:"mode" json:"mode"`               // 0 单机 1 集群
}
```

#### Pair

`Pair`结构体用于SortedSet操作：

```go
type Pair struct {
    Member string  // 成员
    Score  float64 // 分数
}
```

### 类型定义

```go
const (
    RedisNil = goredis.Nil // Redis键不存在错误
)

type (
    Pipeliner = goredis.Pipeliner // 管道操作接口
    Z         = goredis.Z         // SortedSet成员结构
)
```

## 使用示例

### 初始化Redis连接

#### 单机模式

```go
import (
    "github.com/joshqu1985/lego/store/redis"
)

// 初始化Redis单机连接
redisConfig := redis.Config{
    Endpoints: []string{"localhost:6379"},
    Pass:      "password",      // 可选
    Database:  0,               // 数据库索引
    TLS:       false,           // 是否启用TLS
    Mode:      0,               // 0表示单机模式
}

redisStore, err := redis.New(redisConfig)
if err != nil {
    panic(err)
}
```

#### 集群模式

```go
// 初始化Redis集群连接
clusterConfig := redis.Config{
    Endpoints: []string{
        "redis-node1:7001",
        "redis-node2:7002",
        "redis-node3:7003",
    },
    Pass:  "password", // 可选
    TLS:   false,      // 是否启用TLS
    Mode:  1,          // 1表示集群模式
}

redisStore, err := redis.New(clusterConfig)
if err != nil {
    panic(err)
}
```

### String类型操作

```go
import (
    "context"
    "github.com/joshqu1985/lego/store/redis"
)

ctx := context.Background()

// 设置值
err := redisStore.Set(ctx, "key", "value")

// 设置带过期时间的值(秒)
err := redisStore.Set(ctx, "key", "value", 3600) // 1小时过期

// 获取值
var value string
err := redisStore.Get(ctx, "key", &value)

// 检查键是否存在
exists, err := redisStore.Exists(ctx, "key")
if exists == 1 {
    // 键存在
}

// 删除键
delCount, err := redisStore.Del(ctx, "key")

// 设置过期时间
err := redisStore.Expire(ctx, "key", 3600*time.Second)

// 获取剩余过期时间
ttl, err := redisStore.TTL(ctx, "key")

// 自增
val, err := redisStore.IncrBy(ctx, "counter", 1)

// 自减
val, err := redisStore.DecrBy(ctx, "counter", 1)
```

### Hash类型操作

```go
// 设置Hash字段
err := redisStore.HSet(ctx, "user:1", "name", "张三")

// 批量设置Hash字段
fields := map[string]interface{}{
    "name":  "张三",
    "age":   30,
    "email": "zhangsan@example.com",
}
err := redisStore.HMSet(ctx, "user:1", fields)

// 获取Hash字段
var name string
err := redisStore.HGet(ctx, "user:1", "name", &name)

// 批量获取Hash字段
var values struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}
err := redisStore.HMGet(ctx, "user:1", []string{"name", "age", "email"}, &values)

// 获取所有Hash字段的值
var allValues []string
err := redisStore.HVals(ctx, "user:1", &allValues)

// 检查Hash字段是否存在
exists, err := redisStore.HExists(ctx, "user:1", "name")

// 删除Hash字段
delCount, err := redisStore.HDel(ctx, "user:1", "age", "email")

// Hash字段自增
val, err := redisStore.HIncrBy(ctx, "user:1", "visits", 1)

// 获取Hash字段数量
count, err := redisStore.HLen(ctx, "user:1")
```

### SortedSet类型操作

```go
// 添加成员
added, err := redisStore.ZAdd(ctx, "leaderboard", 
    redis.Pair{Member: "user1", Score: 100},
    redis.Pair{Member: "user2", Score: 200},
)

// 获取成员数量
count, err := redisStore.ZCard(ctx, "leaderboard")

// 获取指定分数区间的成员数量
rangeCount, err := redisStore.ZCount(ctx, "leaderboard", 100, 200)

// 获取指定排名范围的成员
pairs, err := redisStore.ZRangeWithScores(ctx, "leaderboard", 0, 9) // 前10名

// 获取指定分数范围的成员
pairs, err := redisStore.ZRangeByScoreWithScores(ctx, "leaderboard", 100, 200)

// 获取成员分数
score, err := redisStore.ZScore(ctx, "leaderboard", "user1")

// 增加成员分数
newScore, err := redisStore.ZIncrBy(ctx, "leaderboard", "user1", 10)

// 删除成员
delCount, err := redisStore.ZRem(ctx, "leaderboard", "user1", "user2")
```

### List类型操作

```go
// 左推元素
len, err := redisStore.LPush(ctx, "list", "value1", "value2")

// 右推元素
len, err := redisStore.Rpush(ctx, "list", "value3", "value4")

// 左弹出元素
var value string
err := redisStore.LPop(ctx, "list", &value)

// 右弹出元素
var value string
err := redisStore.RPop(ctx, "list", &value)

// 获取列表长度
len, err := redisStore.LLen(ctx, "list")
```

### 管道操作

```go
// 使用管道执行多个命令
err := redisStore.Pipelined(ctx, func(pipe redis.Pipeliner) error {
    // 在管道中执行多个命令
    pipe.Set(ctx, "key1", "value1")
    pipe.Set(ctx, "key2", "value2")
    pipe.Incr(ctx, "counter")
    return nil
})
```

### 自定义类型序列化

模块内置了对多种数据类型的序列化支持：

```go
// 保存自定义结构体
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

user := User{Name: "张三", Age: 30, Email: "zhangsan@example.com"}
err := redisStore.Set(ctx, "user:1", user)

// 读取自定义结构体
var loadedUser User
err := redisStore.Get(ctx, "user:1", &loadedUser)
```

## 支持的数据类型

模块的`Format`和`Scan`函数支持以下数据类型的自动序列化和反序列化：

- 基本类型：string, []byte, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool
- 指针类型：上述基本类型的指针
- 时间类型：time.Time, time.Duration
- 网络类型：net.IP
- 复杂类型：struct, map, slice等(通过JSON序列化)

## 最佳实践

1. **连接复用**：创建Redis客户端实例后，应在应用程序生命周期内重用
2. **错误处理**：正确处理Redis操作可能返回的错误，特别是`RedisNil`错误
3. **上下文管理**：合理使用context控制操作超时和取消
4. **连接池配置**：模块内部已配置了基本的连接池参数，如有特殊需求可考虑扩展
5. **键命名规范**：使用命名空间和分隔符组织键名，如`{prefix}:{entity}:{id}`
6. **批量操作**：使用管道(Pipelined)进行批量操作，提高性能
7. **过期时间**：为临时数据设置合理的过期时间，避免内存溢出
8. **集群模式**：在生产环境中考虑使用集群模式提高可用性和性能

## 注意事项

1. **密码安全**：避免在代码中硬编码Redis密码等敏感信息
2. **TLS配置**：在生产环境中，特别是通过公网访问时，建议启用TLS加密
3. **错误处理**：操作返回`RedisNil`表示键不存在，这通常不是错误状态
4. **序列化限制**：复杂类型通过JSON序列化，可能会有性能开销
5. **集群模式限制**：集群模式下，一些操作(如跨槽位事务)可能受限
6. **并发安全**：`Store`结构体本身是并发安全的，但具体的Redis操作需要考虑并发安全

## 依赖说明

- **github.com/redis/go-redis/v9**: Redis客户端库
- **github.com/joshqu1985/lego/encoding/json**: JSON序列化库
