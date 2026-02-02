# Broker 消息队列模块

## 模块概述

Broker模块为Lego框架提供了统一的消息队列抽象层，封装了多种消息中间件的实现，提供一致的接口，使开发者可以轻松切换底层消息队列实现而无需修改业务代码。

## 核心特性

- **统一接口**：提供标准化的生产者和消费者接口
- **多实现支持**：支持内存队列、Redis Stream、Kafka、RocketMQ和Pulsar等多种消息中间件
- **延迟消息**：部分实现支持延迟消息发送和消费
- **配置简单**：通过统一的配置结构配置不同的消息中间件
- **灵活切换**：只需修改配置即可切换底层实现

## 接口定义

### Producer 生产者接口
```go
type Producer interface {
    // Send 发送消息到指定主题
    Send(ctx context.Context, topic string, msg *Message) error
    // SendDelay 发送延迟消息到指定主题，stamp为延迟时间戳（毫秒）
    SendDelay(ctx context.Context, topic string, msg *Message, stamp int64) error
    // Close 关闭生产者连接，释放资源
    Close() error
}
```

### Consumer 消费者接口
```go
type Consumer interface {
    // Register 注册主题的消费回调函数
    Register(topic string, f ConsumeCallback) error
    // Start 启动消费者，开始接收消息
    Start() error
    // Close 关闭消费者连接，释放资源
    Close() error
}
```

### 消息结构
```go
type Message struct {
    Payload    []byte            // 消息体内容
    Key        string            // 消息键，用于分区或去重
    Properties map[string]string // 消息属性，用于扩展信息
    MessageId  string            // 消息ID，发送后自动生成
    Topic      string            // 主题，内部使用
}
```

Message结构体提供了两个便捷方法：
- `SetTag(tag string)`: 设置消息标签
- `GetTag()`: 获取消息标签

### 消费回调函数
```go
type ConsumeCallback func(context.Context, *Message) error
```

## 配置结构

```go
type Config struct {
    Source    string            // 消息队列类型：memory、redis、kafka、rocketmq、pulsar
    Endpoints []string          // 消息队列服务地址列表
    AppId     string            // 应用ID（根据不同实现有不同用途）
    GroupId   string            // 消费组ID
    Topics    map[string]string // 主题映射，key为逻辑主题，value为实际主题
    AccessKey string            // 访问密钥
    SecretKey string            // 密钥
}
```

## 快速开始

### 生产者示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/joshqu1985/lego/broker"
)

func main() {
    // 创建配置
    conf := &broker.Config{
        Source:    "redis",
        Endpoints: []string{"127.0.0.1:6379"},
        GroupId:   "order-service",
        Topics:    map[string]string{"order_created": "orders.created"},
    }
    
    // 创建生产者
    producer, err := broker.NewProducer(conf)
    if err != nil {
        log.Fatalf("创建生产者失败: %v", err)
    }
    defer producer.Close()
    
    // 发送消息
    msg := broker.NewMessage([]byte(`{"order_id":"12345","amount":99.99}`))
    // msg.SetKey("example_1")
    // msg.SetProperties(map[string]string{
    //     "example_2": "new_order",
    // })
    // msg.SetTag("example_tag")
    
    if err := producer.Send(context.Background(), "order_created", msg); err != nil {
        log.Printf("发送消息失败: %v", err)
    } else {
        fmt.Println("消息发送成功")
    }
    
    // 发送延迟消息（10秒后）
    delayTime := time.Now().Add(10 * time.Second).UnixMilli()
    if err := producer.SendDelay(context.Background(), "order_created", msg, delayTime); err != nil {
        log.Printf("发送延迟消息失败: %v", err)
    }
}
```

### 消费者示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/joshqu1985/lego/broker"
)

func main() {
    // 创建配置
    conf := &broker.Config{
        Source:    "redis",
        Endpoints: []string{"127.0.0.1:6379"},
        GroupId:   "payment-service",
        Topics:    map[string]string{"order_created": "orders.created"},
    }
    
    // 创建消费者
    consumer, err := broker.NewConsumer(conf)
    if err != nil {
        log.Fatalf("创建消费者失败: %v", err)
    }
    defer consumer.Close()
    
    // 注册消息处理函数
    err = consumer.Register("order_created", func(ctx context.Context, msg *broker.Message) error {
        fmt.Printf("收到消息: %s\n", string(msg.GetPayload()))
        fmt.Printf("消息ID: %s\n", msg.GetMsgId())
        fmt.Printf("消息键: %s\n", msg.GetKey())
        fmt.Printf("消息类型: %s\n", msg.GetProperty("type"))
        // 处理消息逻辑...
        return nil // 返回nil表示消息处理成功
    })
    if err != nil {
        log.Fatalf("注册消费者失败: %v", err)
    }
    
    // 启动消费者（在goroutine中运行，避免阻塞）
    go func() {
        if err := consumer.Start(); err != nil {
            log.Printf("消费者启动失败: %v", err)
        }
    }()
    
    // 保持程序运行
    fmt.Println("消费者已启动，按Ctrl+C退出...")
    select {}
}
```

## 各实现详细说明

### 1. 内存队列（Memory）

- **特点**：基于链表和优先级队列实现，纯内存操作，无外部依赖
- **适用场景**：测试环境、单机应用、对消息可靠性要求不高的场景
- **限制**：
  - 进程重启后消息丢失
  - 每个主题最多10000条消息，超出后Send操作会失败
- **支持特性**：
  - ✅ 延迟消息（通过优先级队列实现）
  - ❌ 持久化
  - ❌ 高可用

```go
conf := &broker.Config{
    Source:  "memory",
    GroupId: "test-group",
    Topics:  map[string]string{"topic1": "real_topic1"},
}
```

### 2. Redis Stream

- **特点**：基于Redis Stream实现，轻量级、易部署
- **版本要求**：Redis 5.0及以上
- **限制**：每个主题最多10000条消息，超出后Send操作会失败
- **支持特性**：
  - ✅ 延迟消息（使用ZSET和HASH实现）
  - ✅ 基本持久化
  - ⚠️ 有限的高可用（依赖Redis集群）

```go
conf := &broker.Config{
    Source:    "redis",
    Endpoints: []string{"127.0.0.1:6379"},
    GroupId:   "order-service",
    Topics:    map[string]string{"order_created": "orders.created"},
    // 可选的Redis认证信息
    // AccessKey: "username",
    // SecretKey: "password",
}
```

### 3. Kafka

- **特点**：基于Apache Kafka实现，高吞吐量、高可靠性
- **底层依赖**：github.com/segmentio/kafka-go
- **支持特性**：
  - ❌ 延迟消息
  - ✅ 持久化
  - ✅ 高可用
  - ✅ 高吞吐量

```go
conf := &broker.Config{
    Source:    "kafka",
    Endpoints: []string{"127.0.0.1:9092", "127.0.0.1:9093"},
    GroupId:   "user-service",
    Topics:    map[string]string{"user_registered": "users.registered"},
    // 可选的Kafka认证信息
    // AccessKey: "username",
    // SecretKey: "password",
}
```

### 4. RocketMQ

- **特点**：基于Apache RocketMQ实现，功能丰富、支持广泛
- **底层依赖**：github.com/apache/rocketmq-clients/golang/v5
- **支持特性**：
  - ✅ 延迟消息
  - ✅ 持久化
  - ✅ 高可用
  - ✅ 丰富的消息类型

```go
conf := &broker.Config{
    Source:    "rocketmq",
    Endpoints: []string{"127.0.0.1:8081"},
    GroupId:   "notification-service",
    Topics:    map[string]string{"send_sms": "notifications.sms"},
    // RocketMQ认证信息
    // AccessKey: "your-access-key",
    // SecretKey: "your-secret-key",
}
```

### 5. Pulsar

- **特点**：基于Apache Pulsar实现，云原生设计、多租户支持
- **底层依赖**：github.com/apache/pulsar-client-go/pulsar
- **支持特性**：
  - ✅ 延迟消息
  - ✅ 持久化
  - ✅ 高可用
  - ✅ 多租户

```go
conf := &broker.Config{
    Source:    "pulsar",
    Endpoints: []string{"pulsar://localhost:6650"},
    GroupId:   "analytics-service",
    Topics:    map[string]string{
        "user_activity": "persistent://public/default/user-activity",
    },
    // Pulsar认证信息
    // AccessKey: "your-access-key",
    // SecretKey: "your-secret-key",
}
```

## 最佳实践

1. **错误处理**：始终检查Send、SendDelay和Start方法返回的错误
2. **资源释放**：使用defer确保在程序退出时调用Close方法释放资源
3. **消费者设计**：
   - 消息处理函数应该幂等
   - 避免在回调中执行长时间运行的操作
   - 合理设置超时时间
4. **主题命名**：使用规范的主题命名，如`domain.event_type`
5. **配置管理**：结合Lego的configor模块进行配置管理

## 注意事项

1. 内存队列仅适用于测试或单机应用，生产环境请使用Redis、Kafka等持久化方案
2. 延迟消息在不同实现中的行为可能有所差异
3. Redis实现和内存实现有消息数量限制，请注意监控队列大小
4. 生产环境使用时请确保正确配置认证信息和网络访问权限
5. 消费者启动后需要保持运行，建议在独立的goroutine中启动并处理错误
