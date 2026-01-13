# MongoDB 存储模块

**MongoDB 存储模块** 提供了对MongoDB数据库的统一封装，简化了数据库连接管理、操作执行和事务处理，同时集成了性能监控功能。

## 核心特性

- **统一接口**：提供简洁的接口执行MongoDB操作，避免直接引用MongoDB SDK
- **连接池管理**：支持配置最大/最小连接池大小和连接生命周期
- **事务支持**：内置事务处理机制，简化事务操作
- **超时控制**：自动设置操作超时时间，避免长时间阻塞
- **性能监控**：集成metrics指标监控，记录执行时间和慢查询统计
- **灵活配置**：支持丰富的连接配置选项

## 接口定义

```go
// Store MongoDB存储客户端
type Store struct {
    client *mongo.Client
    db     *mongo.Database
}

// Exec 执行MongoDB操作，自动处理上下文和监控
func (this *Store) Exec(fn func(context.Context, *DB) error, labels ...string) error

// Transaction 执行MongoDB事务操作，自动处理事务生命周期和监控
func (this *Store) Transaction(fn func(context.Context, *DB) error, labels ...string) error

// DB 类型别名，避免业务代码直接引用mongo sdk
type DB = mongo.Database
```

## 配置选项

```go
// Config MongoDB配置
type Config struct {
    Endpoint    string // 数据库连接地址，如: 10.X.X.X:27017,10.X.X.X:27017/admin
    User        string // 用户名
    Pass        string // 密码
    Opts        string // 连接选项，如: replicaSet=cmgo-*******&authSource=admin
    Database    string // 数据库名称
    MaxPoolSize uint64 // 最大连接池大小
    MinPoolSize uint64 // 最小连接池大小
    MaxLife     uint64 // 连接最大生命周期（毫秒）
}
```

## 初始化

```go
import (
    "github.com/joshqu1985/lego/store/mongo"
)

// 创建配置
conf := &mongo.Config{
    Endpoint:    "127.0.0.1:27017",
    User:        "admin",
    Pass:        "password",
    Opts:        "replicaSet=rs0&authSource=admin",
    Database:    "mydatabase",
    MaxPoolSize: 100,
    MinPoolSize: 10,
    MaxLife:     60000, // 60秒
}

// 初始化MongoDB存储
store, err := mongo.New(conf)
if err != nil {
    panic(err)
}
```

## 使用示例

### 基本查询操作

```go
import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/v2/bson"
)

// 定义文档结构
type User struct {
    ID       string `bson:"_id"`
    Name     string `bson:"name"`
    Age      int    `bson:"age"`
    Email    string `bson:"email"`
}

// 执行查询
var user User
err := store.Exec(func(ctx context.Context, db *mongo.DB) error {
    collection := db.Collection("users")
    filter := bson.M{"_id": "user123"}
    return collection.FindOne(ctx, filter).Decode(&user)
}, "find_user") // 监控标签

if err != nil {
    fmt.Printf("查询失败: %v\n", err)
} else {
    fmt.Printf("找到用户: %s\n", user.Name)
}
```

### 插入操作

```go
// 执行插入
newUser := User{
    ID:    "user456",
    Name:  "张三",
    Age:   30,
    Email: "zhangsan@example.com",
}

err := store.Exec(func(ctx context.Context, db *mongo.DB) error {
    collection := db.Collection("users")
    _, err := collection.InsertOne(ctx, newUser)
    return err
}, "insert_user")

if err != nil {
    fmt.Printf("插入失败: %v\n", err)
} else {
    fmt.Println("插入成功")
}
```

### 更新操作

```go
// 执行更新
err := store.Exec(func(ctx context.Context, db *mongo.DB) error {
    collection := db.Collection("users")
    filter := bson.M{"_id": "user123"}
    update := bson.M{"$set": bson.M{"age": 31}}
    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if result.ModifiedCount == 0 {
        return fmt.Errorf("用户不存在")
    }
    return nil
}, "update_user")

if err != nil {
    fmt.Printf("更新失败: %v\n", err)
} else {
    fmt.Println("更新成功")
}
```

### 删除操作

```go
// 执行删除
err := store.Exec(func(ctx context.Context, db *mongo.DB) error {
    collection := db.Collection("users")
    filter := bson.M{"_id": "user123"}
    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return fmt.Errorf("用户不存在")
    }
    return nil
}, "delete_user")

if err != nil {
    fmt.Printf("删除失败: %v\n", err)
} else {
    fmt.Println("删除成功")
}
```

### 事务操作

```go
// 执行事务操作
err := store.Transaction(func(ctx context.Context, db *mongo.DB) error {
    // 获取集合
    usersCollection := db.Collection("users")
    ordersCollection := db.Collection("orders")
    
    // 步骤1: 更新用户余额
    filter := bson.M{"_id": "user123"}
    update := bson.M{"$inc": bson.M{"balance": -100}}
    result, err := usersCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if result.ModifiedCount == 0 {
        return fmt.Errorf("用户不存在或余额不足")
    }
    
    // 步骤2: 创建订单
    order := map[string]interface{}{
        "_id":      "order789",
        "user_id":  "user123",
        "amount":   100,
        "status":   "pending",
        "created_at": time.Now(),
    }
    _, err = ordersCollection.InsertOne(ctx, order)
    if err != nil {
        return err
    }
    
    return nil
}, "create_order")

if err != nil {
    fmt.Printf("事务执行失败: %v\n", err) // 事务会自动回滚
} else {
    fmt.Println("事务执行成功")
}
```

## 性能监控

模块自动集成了性能监控指标，记录以下信息：

1. **执行时间**：通过`mongo_client_duration_ms`直方图记录操作耗时
2. **慢查询统计**：通过`mongo_client_slow_count`计数器记录慢查询次数（>500ms）
3. **标签支持**：可以通过`labels`参数为操作添加自定义标签，方便分类统计

监控指标分桶定义：[3ms, 5ms, 10ms, 50ms, 100ms, 250ms, 500ms, 1000ms, 2000ms, 5000ms]

## 最佳实践

1. **连接池配置**：根据应用规模合理设置MaxPoolSize和MinPoolSize，避免连接过多或过少
2. **超时控制**：虽然模块内置了10秒超时，但对于特定操作可以在回调函数中使用自定义超时
3. **错误处理**：始终检查操作返回的错误，特别是事务操作
4. **监控标签**：使用有意义的监控标签，方便后续分析性能瓶颈
5. **事务使用**：只在需要原子性的场景下使用事务，避免过度使用影响性能
6. **索引优化**：确保查询字段有适当的索引，提高查询性能
7. **批量操作**：对于大量数据操作，考虑使用批量插入、更新等操作

## 注意事项

1. 操作超时时间默认设置为10秒，可根据实际需求在回调函数中调整
2. 慢查询阈值默认设置为500ms，可以根据应用性能要求进行调整
3. 事务操作需要MongoDB 4.0以上版本支持
4. 连接字符串中的特殊字符需要进行URL编码
5. 在高并发场景下，建议增加监控检查数据库连接状态
6. 对于分片集群，需要正确配置连接字符串中的副本集信息
7. 定期检查和优化数据库索引，确保查询性能
