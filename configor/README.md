# Configor 配置管理模块

## 模块概述

Configor模块为Lego框架提供了统一的配置中心客户端封装库，支持多种配置源（本地文件、Apollo、Nacos、etcd），提供一致的接口，使开发者可以轻松切换不同的配置管理方案。

## 核心特性

- **统一接口**：提供标准化的配置加载和变更监听接口
- **多源支持**：支持本地文件、Apollo、Nacos、etcd等多种配置源
- **配置热更新**：支持配置变更的实时监听和通知
- **灵活编码**：支持TOML、YAML、JSON等多种配置格式
- **无缝切换**：只需修改配置文件即可切换底层配置源实现

## 接口定义

### Configor 接口
```go
type Configor interface {
    // Load 将配置加载到指定结构体中
    Load(v any) error
}
```

### ChangeSet 结构体
```go
type ChangeSet struct {
    Timestamp time.Time // 配置变更时间戳
    Value     []byte    // 配置内容
}
```

### SourceConfig 结构体
```go
type SourceConfig struct {
    Source    string   // 配置源类型：local、apollo、nacos、etcd
    Endpoints []string // 配置源服务地址列表
    Cluster   string   // 集群标识（各配置源含义不同）
    AppId     string   // 应用ID（各配置源含义不同）
    Namespace string   // 命名空间（主要用于Apollo）
    AccessKey string   // 访问密钥
    SecretKey string   // 密钥
}
```

各配置源中字段的具体含义：
- **Apollo**: 
  - Cluster: 应用所属的集群（默认default）
  - AppId: 用来标识应用身份（服务名，唯一）
  - Namespace: 配置项的集合
- **Nacos**: 
  - Cluster: Nacos的namespaceid
  - AppId: Nacos的dataId
- **Etcd**: 
  - Cluster: etcd的key前缀
  - AppId: etcd的key前缀

## 配置选项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| WithToml() | 使用TOML格式解析配置 | 是 |
| WithYaml() | 使用YAML格式解析配置 | 否 |
| WithJson() | 使用JSON格式解析配置 | 否 |
| WithWatch(f ChangeNotify) | 设置配置变更通知回调函数 | 无 |

## 使用指南

### 基本用法

1. 创建配置源配置文件（例如：config.toml）

```toml
# 本地文件配置示例
source = "local"

# 或 Apollo配置示例
# source = "apollo"
# endpoints = ["http://127.0.0.1:8080"]
# cluster = "default"
# app_id = "your-app-id"
# namespace = "application"
# access_key = ""
# secret_key = "your-secret"

# 或 Nacos配置示例
# source = "nacos"
# endpoints = ["127.0.0.1:8848"]
# cluster = "public"
# app_id = "your-data-id"
# access_key = "your-access-key"
# secret_key = "your-secret-key"

# 或 Etcd配置示例
# source = "etcd"
# endpoints = ["127.0.0.1:2379"]
# cluster = "default"
# app_id = "your-app"
# access_key = "username"
# secret_key = "password"
```

2. 定义配置结构体

```go
type AppConfig struct {
    Server struct {
        Host string `toml:"host" yaml:"host" json:"host"`
        Port int    `toml:"port" yaml:"port" json:"port"`
    } `toml:"server" yaml:"server" json:"server"`
    
    Database struct {
        DSN string `toml:"dsn" yaml:"dsn" json:"dsn"`
    } `toml:"database" yaml:"database" json:"database"`
}
```

3. 初始化配置器并加载配置

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/joshqu1985/lego/configor"
)

func main() {
    // 创建配置器（默认使用TOML格式）
    cfg, err := configor.New("config.toml")
    if err != nil {
        log.Fatalf("创建配置器失败: %v", err)
    }
    
    // 定义配置结构体
    var appConfig AppConfig
    
    // 加载配置
    if err := cfg.Load(&appConfig); err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }
    
    fmt.Printf("服务器配置: %s:%d\n", appConfig.Server.Host, appConfig.Server.Port)
}
```

### 配置变更监听

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/joshqu1985/lego/configor"
)

func main() {
    // 创建配置器并设置变更监听
    cfg, err := configor.New("config.toml", 
        configor.WithWatch(func(set configor.ChangeSet) {
            fmt.Printf("配置已更新，时间: %v\n", set.Timestamp)
            // 这里可以重新加载配置或执行其他操作
            var appConfig AppConfig
            if err := cfg.Load(&appConfig); err != nil {
                log.Printf("重新加载配置失败: %v", err)
            } else {
                fmt.Printf("新的服务器配置: %s:%d\n", appConfig.Server.Host, appConfig.Server.Port)
            }
        }),
    )
    if err != nil {
        log.Fatalf("创建配置器失败: %v", err)
    }
    
    // 初始加载配置
    var appConfig AppConfig
    if err := cfg.Load(&appConfig); err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }
    
    fmt.Println("配置监听已启动，按Ctrl+C退出...")
    select {} // 保持程序运行
}
```

### 选择不同的配置格式

```go
// 使用YAML格式
cfg, err := configor.New("config.yaml", configor.WithYaml())

// 使用JSON格式
cfg, err := configor.New("config.json", configor.WithJson())
```

## 各配置源详解

### 1. 本地文件

- **特点**：基于文件系统的配置管理，简单易用
- **适用场景**：单机应用、开发环境、配置变更频率低的场景
- **实现机制**：使用fsnotify监听文件变化
- **配置示例**：
  ```toml
  source = "local"
  ```

### 2. Apollo

- **特点**：携程开源的配置中心，功能丰富，支持灰度发布
- **适用场景**：微服务架构、需要动态配置管理的生产环境
- **依赖**：github.com/apolloconfig/agollo/v4
- **配置示例**：
  ```toml
  source = "apollo"
  endpoints = ["http://127.0.0.1:8080"]
  cluster = "default"
  app_id = "your-app-id"
  namespace = "application"
  secret_key = "your-secret"
  ```

### 3. Nacos

- **特点**：阿里巴巴开源的服务发现和配置管理平台
- **适用场景**：微服务架构、云原生应用
- **配置示例**：
  ```toml
  source = "nacos"
  endpoints = ["127.0.0.1:8848"]
  cluster = "public"
  app_id = "your-data-id"
  access_key = "your-access-key"
  secret_key = "your-secret-key"
  ```

### 4. etcd

- **特点**：分布式键值存储，高可用，强一致性
- **适用场景**：分布式系统、高可用架构
- **依赖**：go.etcd.io/etcd/client/v3
- **配置示例**：
  ```toml
  source = "etcd"
  endpoints = ["127.0.0.1:2379", "127.0.0.1:2380"]
  cluster = "default"
  app_id = "your-app"
  access_key = "username"
  secret_key = "password"
  ```

## 最佳实践

1. **配置分离**：将配置源配置与应用配置分离，便于环境切换
2. **默认值处理**：为配置结构体字段设置合理的默认值
3. **配置验证**：加载配置后进行必要的验证，确保配置有效性
4. **错误处理**：妥善处理配置加载和监听过程中的错误
5. **安全考虑**：敏感配置信息（如密码、密钥）避免明文存储

## 注意事项

1. 本地文件配置在服务重启时会重新加载，无需额外处理
2. Apollo、Nacos、etcd配置源在配置变更时会自动通知应用
3. 使用配置变更监听时，需注意处理并发更新场景
4. 生产环境建议使用分布式配置中心（Apollo、Nacos或etcd），便于集中管理和动态更新
5. 配置变更可能影响系统行为，建议在变更时采取灰度策略
