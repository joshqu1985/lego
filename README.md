# Lego 框架

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/joshqu1985/lego)
![GitHub](https://img.shields.io/github/license/joshqu1985/lego)

## 项目概述

Lego是一个功能丰富、高性能的Go语言微服务框架，提供了完整的微服务开发基础设施。框架采用模块化设计，包含配置管理、消息队列、RPC通信、数据存储等核心组件，旨在简化微服务开发流程，提高开发效率和系统稳定性。

## 核心特性

- **模块化设计**：各功能组件独立封装，可单独使用也可组合使用
- **统一接口**：提供一致的API设计，降低学习成本
- **多实现支持**：关键组件支持多种底层实现，可灵活切换
- **企业级特性**：内置服务注册发现、熔断器、监控等企业级功能
- **高性能**：优化的实现和精心设计的API，确保系统高效运行
- **易于扩展**：提供清晰的接口和扩展点，方便自定义实现

## 架构概览

Lego框架由多个核心模块组成，每个模块负责特定的功能域：

```
lego/
├── broker/      # 消息队列模块，支持多种消息中间件
├── configor/    # 配置管理模块，支持多种配置源
├── container/   # 容器和数据结构模块
├── encoding/    # 编码解码模块
├── logs/        # 日志模块
├── metrics/     # 指标收集模块
├── store/       # 数据存储模块
├── transport/   # 通信传输模块
└── utils/       # 通用工具库
```

### 模块关系图

```
┌───────────────┐      ┌───────────────┐      ┌───────────────┐
│   configor    │      │     logs      │      │    metrics    │
│  (配置管理)   │      │   (日志系统)  │      │  (指标监控)   │
└───────┬───────┘      └───────┬───────┘      └───────┬───────┘
        │                      │                      │
        └────────────┬─────────┴───────────┬──────────┘
                     │                     │
        ┌────────────▼────────────┐ ┌──────▼────────────┐
        │         broker         │ │      transport    │
        │      (消息队列)        │ │    (通信传输)     │
        └────────────┬────────────┘ └──────┬────────────┘
                     │                     │
                     └─────────┬───────────┘
                               │
                       ┌───────▼───────┐
                       │      store    │
                       │    (数据存储)  │
                       └───────┬───────┘
                               │
                       ┌───────▼───────┐
                       │     utils     │
                       │   (工具库)    │
                       └───────────────┘
```

## 核心模块介绍

### 1. broker - 消息队列模块

统一的消息队列抽象层，封装了多种消息中间件的实现：

- 支持内存队列、Redis Stream、Kafka、RocketMQ、Pulsar
- 提供标准化的生产者和消费者接口
- 支持延迟消息发送和消费
- 配置简单，灵活切换底层实现
- 内置错误常量定义，便于错误处理

**支持的常量**：
- `SourceKafka`、`SourceRocket`、`SourcePulsar`、`SourceRedis`、`SourceMemory`
- `ErrEndpointsEmpty`、`ErrSubscriberNil`、`ErrMessageIsNil`、`ErrQueueIsFull`、`ErrTopicNotFound`

**核心接口**：
- `Producer`：生产者接口，提供 `Send`、`SendDelay`、`Close` 方法
- `Consumer`：消费者接口，提供 `Register`、`Start`、`Close` 方法

### 2. configor - 配置管理模块

统一的配置中心客户端封装库：

- 支持本地文件、Apollo、Nacos、etcd等多种配置源
- 提供配置热更新和变更监听能力
- 支持TOML、YAML、JSON等多种配置格式
- 无缝切换不同的配置管理方案
- 简化的API设计，自动检测配置源

**支持的常量**：
- `ENCODING_JSON`、`ENCODING_YAML`、`ENCODING_TOML`
- `SOURCE_ETCD`、`SOURCE_APOLLO`、`SOURCE_NACOS`

**核心接口**：
- `Configor`：配置管理接口，提供 `Load` 方法
- `SourceConfig`：配置源配置结构，支持多种配置中心参数

### 3. transport - 通信传输模块

高性能的RPC和REST通信组件：

- **rpc子模块**：基于gRPC实现的RPC框架
  - 集成服务注册发现、熔断器、指标监控
  - 支持超时控制、异常恢复、优雅关闭
  - 内置健康检查和负载均衡
  - 支持自定义拦截器（Unary和Stream）
  - 内置监控和恢复拦截器

- **rest子模块**：RESTful API框架
  - 基于Gin实现的HTTP服务
  - 提供中间件机制和响应格式化
  - 支持优雅关闭和信号处理
  - 内置健康检查、pprof和指标端点
  - 支持服务注册发现

**核心特性**：
- 统一的服务注册发现接口（Nacos、etcd）
- 内置监控和日志集成
- 支持TLS和自定义配置

### 4. store - 数据存储模块

统一的数据存储抽象：

- 支持MongoDB、Redis等NoSQL数据库
- 支持关系型数据库ORM（基于GORM）
- 支持多种对象存储（S3、OSS、COS、OBS）

**Redis特性**：
- 支持单机和集群模式
- 内置连接超时控制（5秒）
- 支持TLS加密连接
- 提供完整的Redis客户端接口
- 自动重试机制（最多3次）
- 最小空闲连接数配置（8个）

### 5. logs - 日志模块

灵活的日志记录系统：

- 支持多种日志级别和格式化
- 提供日志滚动和分割功能
- 集成zap高性能日志库

### 6. metrics - 指标收集模块

指标监控和收集功能：

- 集成Prometheus客户端
- 提供丰富的指标收集API
- 支持自定义指标

### 7. utils - 工具库

实用工具集合：

- 加密工具（crypto）
- IP地址处理（ip）
- 分页令牌（pagetoken）
- 协程管理（routine）
- 切片操作（slice）
- 字符串处理（strings）
- 时间处理（utime）

### 8. container - 容器和数据结构

提供各种数据结构和容器实现：

- 批处理器（batcher）
- 位图（bitmap）
- 布隆过滤器（bloom_filter）
- 链表（linked）
- 优先队列（priority）
- 滑动窗口（sliding_window）

### 9. encoding - 编码解码

支持多种数据格式的编解码：

- JSON
- YAML
- TOML

## 最新更新

### v1.x - 代码优化与改进

本次更新对框架进行了全面的代码优化和改进，主要变化包括：

#### API 改进
- **broker 模块**：优化了 Producer 和 Consumer 接口设计，参数改为指针传递，添加了 Close 方法
- **configor 模块**：简化了 New 函数签名，自动检测配置源类型，支持 Option 模式配置
- **transport 模块**：改进了 Server 创建方式，使用 Option 模式，支持服务注册发现集成
- **store 模块**：Redis 客户端添加了连接超时控制，支持 TLS 加密连接

#### 错误处理优化
- 将 panic 改为返回 error，提高代码健壮性
- 添加了统一的错误常量定义（如 `ErrEndpointsEmpty`、`ErrSubscriberNil` 等）
- 改进了错误消息的可读性

#### 代码质量提升
- 统一了代码风格，使用 type () 组合类型定义
- 添加了详细的常量定义，避免硬编码字符串
- 改进了命名规范（如 `this` → `s`）
- 添加了更多的注释和文档

#### 性能优化
- Redis 客户端添加了连接超时控制（5秒）
- 优化了连接池配置（最小空闲连接数 8，最大重试次数 3）
- 改进了并发控制，添加了 sync.RWMutex 保护

#### 架构改进
- 服务注册发现模块（naming）支持 Nacos、etcd 等多种实现
- RPC 和 REST 服务都支持优雅关闭
- 内置了健康检查、pprof 和指标监控端点
- 支持自定义拦截器（Unary 和 Stream）

#### 新增特性
- 支持延迟消息发送（broker.SendDelay）
- 支持配置热更新监听（configor）
- 支持服务注册发现（transport/naming）
- 支持 TLS 加密连接（store/redis）

## 快速开始

### 环境要求

- Go 1.22+（基于go.mod文件）

### 安装

```bash
go get -u github.com/joshqu1985/lego
```

### 基本使用示例

#### 1. 配置管理示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/configor"
)

// AppConfig 应用配置结构
type AppConfig struct {
    Server struct {
        Port int    `json:"port" yaml:"port" toml:"port"`
        Host string `json:"host" yaml:"host" toml:"host"`
    } `json:"server" yaml:"server" toml:"server"`
    Database struct {
        URL string `json:"url" yaml:"url" toml:"url"`
    } `json:"database" yaml:"database" toml:"database"`
}

func main() {
    // 创建配置管理器（支持本地文件、Apollo、Nacos、etcd）
    cfg, err := configor.New("config.toml")
    if err != nil {
        panic(err)
    }
    
    // 加载配置
    var appConfig AppConfig
    if err := cfg.Load(&appConfig); err != nil {
        panic(err)
    }
    
    fmt.Printf("Server running at %s:%d\n", appConfig.Server.Host, appConfig.Server.Port)
}
```

#### 2. 消息队列示例

```go
package main

import (
    "context"
    "fmt"
    "github.com/joshqu1985/lego/broker"
)

func main() {
    // 创建生产者（支持 Kafka、RocketMQ、Pulsar、Redis、Memory）
    producer, err := broker.NewProducer(&broker.Config{
        Source:    broker.SourceKafka,
        Endpoints: []string{"localhost:9092"},
        Topics:    map[string]string{"test-topic": "test-topic"},
    })
    if err != nil {
        panic(err)
    }
    defer producer.Close()

    // 发送消息
    msg := &broker.Message{
        Key:     "test-key",
        Payload: []byte("hello world"),
    }
    if err := producer.Send(context.Background(), "test-topic", msg); err != nil {
        panic(err)
    }

    // 创建消费者
    consumer, err := broker.NewConsumer(&broker.Config{
        Source:    broker.SourceKafka,
        Endpoints: []string{"localhost:9092"},
        Topics:    map[string]string{"test-topic": "test-topic"},
        GroupId:   "test-group",
    })
    if err != nil {
        panic(err)
    }
    defer consumer.Close()

    // 注册消费回调
    if err := consumer.Register("test-topic", func(ctx context.Context, msg *broker.Message) error {
        fmt.Printf("Received message: %s\n", string(msg.Payload))
        return nil
    }); err != nil {
        panic(err)
    }

    // 启动消费者
    if err := consumer.Start(); err != nil {
        panic(err)
    }
}
```

#### 3. RPC服务示例

```go
package main

import (
    "context"
    "log"
    "github.com/joshqu1985/lego/transport/rpc"
    "github.com/joshqu1985/lego/transport/naming"
    pb "path/to/your/proto"
)

// Server 实现服务接口
type Server struct{}

// SayHello 实现SayHello方法
func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
    // 创建服务注册发现（支持 Nacos、etcd）
    naming, err := naming.New(&naming.Config{
        Source:    naming.SOURCE_NACOS,
        Endpoints: []string{"localhost:8848"},
    })
    if err != nil {
        log.Fatalf("Failed to create naming: %v", err)
    }

    // 创建gRPC服务器
    server, err := rpc.NewServer("hello-service", "0.0.0.0:8080",
        rpc.WithNaming(naming),
        rpc.WithRouters(func(s *grpc.Server) {
            pb.RegisterGreeterServer(s, &Server{})
        }),
    )
    if err != nil {
        log.Fatalf("Failed to create server: %v", err)
    }
    
    // 启动服务
    if err := server.Start(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

#### 4. REST服务示例

```go
package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joshqu1985/lego/transport/rest"
    "github.com/joshqu1985/lego/transport/naming"
)

func main() {
    // 创建服务注册发现
    naming, err := naming.New(&naming.Config{
        Source:    naming.SOURCE_NACOS,
        Endpoints: []string{"localhost:8848"},
    })
    if err != nil {
        log.Fatalf("Failed to create naming: %v", err)
    }

    // 创建REST服务器
    server, err := rest.NewServer("api-service", "0.0.0.0:8080",
        rest.WithNaming(naming),
        rest.WithRouters(func(router *rest.Router) {
            router.GET("/health", func(c *rest.Context) {
                c.JSON(200, gin.H{"status": "ok"})
            })
        }),
    )
    if err != nil {
        log.Fatalf("Failed to create server: %v", err)
    }
    
    // 启动服务
    if err := server.Start(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

#### 5. Redis存储示例

```go
package main

import (
    "context"
    "log"
    "github.com/joshqu1985/lego/store/redis"
)

func main() {
    // 创建Redis客户端（支持单机和集群模式）
    store, err := redis.New(redis.Config{
        Endpoints: []string{"localhost:6379"},
        Pass:      "password",
        Database:  0,
        Mode:      0, // 0 单机 1 集群
        TLS:       false,
    })
    if err != nil {
        log.Fatalf("Failed to create redis store: %v", err)
    }

    // 使用Redis
    ctx := context.Background()
    if err := store.Set(ctx, "key", "value", 0); err != nil {
        log.Fatalf("Failed to set key: %v", err)
    }

    var val string
    if err := store.Get(ctx, "key", &val); err != nil {
        log.Fatalf("Failed to get key: %v", err)
    }

    log.Printf("Got value: %s", val)
}
```

## 最佳实践

1. **模块化使用**：根据项目需求选择性使用框架模块，避免引入不必要的依赖
2. **配置集中管理**：使用configor模块进行统一的配置管理，支持环境隔离和动态更新
3. **消息队列解耦**：通过broker模块实现服务间的异步通信，提高系统弹性
4. **合理使用缓存**：利用Redis客户端提供的缓存功能，优化系统性能
5. **统一日志格式**：使用logs模块保证日志格式一致性，便于日志分析和问题排查
6. **监控指标收集**：集成metrics模块，收集关键业务指标，实现可观测性
7. **服务注册发现**：使用naming模块进行服务注册和发现，支持Nacos、etcd等多种实现
8. **优雅关闭**：所有服务端都支持优雅关闭，确保资源正确释放
9. **错误处理**：使用框架提供的错误常量，统一错误处理逻辑
10. **连接管理**：Redis等存储组件内置连接超时和重试机制，提高系统稳定性

## 版本兼容性

- Go 1.22+（基于go.mod文件）
- 向后兼容，API稳定

## 许可证

[MIT License](https://github.com/joshqu1985/lego/blob/main/LICENSE)