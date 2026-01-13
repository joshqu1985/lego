# 服务注册与发现模块

## 概述

`transport/naming` 模块提供了统一的服务注册与发现接口，支持多种后端实现（如 etcd、Nacos），用于微服务架构中的服务治理。该模块负责服务实例的注册、注销、发现以及服务变更的实时监听，是构建分布式系统的重要基础设施。

## 核心特性

- 统一的服务注册与发现接口
- 多后端支持（etcd、Nacos、Pass-through）
- 服务实例的注册与注销
- 实时服务变更监听
- 健康检查与自动剔除
- 租约机制（etcd实现）
- 多集群支持

## 接口定义

### Naming 接口

```go
type Naming interface {
    Endpoints() []string               // 获取命名服务的端点列表
    Name() string                      // 获取命名服务的名称
    Register(key, val string) error     // 注册服务实例
    Deregister(key string) error        // 注销服务实例
    Service(key string) RegService      // 获取服务对象
}
```

### RegService 接口

```go
type RegService interface {
    Name() string                   // 获取服务名称
    Addrs() ([]string, error)       // 获取服务实例地址列表
    AddListener(func())             // 添加服务变更监听器
}
```

## 配置结构

```go
type Config struct {
    Source    string   // 命名服务类型: "nacos" | "etcd"
    Endpoints []string // 命名服务端点列表
    Cluster   string   // 集群标识: nacos namespaceid | etcd key前缀
    AccessKey string   // 访问密钥: nacos access_key | etcd username
    SecretKey string   // 密钥: nacos secret_key | etcd password
}
```

## 实现类型

### 1. etcd 实现

基于 etcd 的服务注册与发现实现，主要特点：
- 使用 etcd 租约机制保持服务活跃性
- 支持前缀匹配的服务发现
- 实时监听服务变更
- 自动续约机制

### 2. Nacos 实现

基于 Nacos 的服务注册与发现实现，主要特点：
- 支持 Nacos 服务注册与发现
- 健康检查机制
- 实例权重设置
- 实时服务订阅

### 3. Pass-through 实现

一个简单的直通实现，主要用于测试或不需要实际服务注册发现的场景：
- 不执行实际的服务注册
- 返回简单的服务地址
- 无实际监听功能

## 使用示例

### 创建命名服务实例

```go
import "github.com/joshqu1985/lego/transport/naming"

// 创建 etcd 命名服务配置
config := &naming.Config{
    Source:    "etcd",
    Endpoints: []string{"localhost:2379"},
    Cluster:   "development",
    AccessKey: "user",
    SecretKey: "password",
}

// 初始化命名服务
namingClient, err := naming.New(config)
if err != nil {
    log.Fatalf("Failed to initialize naming service: %v", err)
}
```

### 注册服务实例

```go
// 注册服务
// key: 服务名称
// val: 服务实例地址 (host:port)
err = namingClient.Register("my-service", "192.168.1.100:8080")
if err != nil {
    log.Printf("Failed to register service: %v", err)
}
```

### 发现服务实例

```go
// 获取服务对象
service := namingClient.Service("my-service")

// 获取服务实例地址列表
addrs, err := service.Addrs()
if err != nil {
    log.Printf("Failed to get service addresses: %v", err)
}

// 打印服务地址
for _, addr := range addrs {
    log.Printf("Service address: %s", addr)
}
```

### 监听服务变更

```go
// 添加服务变更监听器
service.AddListener(func() {
    // 服务发生变更时执行此回调
    addrs, err := service.Addrs()
    if err != nil {
        log.Printf("Failed to get updated service addresses: %v", err)
        return
    }
    log.Printf("Service addresses updated: %v", addrs)
})
```

### 注销服务实例

```go
// 注销服务实例
err = namingClient.Deregister("my-service")
if err != nil {
    log.Printf("Failed to deregister service: %v", err)
}
```

## 高级功能

### 集群选择

通过配置 `Cluster` 参数可以指定服务注册与发现的集群环境：

```go
// 生产环境配置
prodConfig := &naming.Config{
    Source:  "etcd",
    Cluster: "production",
    // 其他配置...
}

// 测试环境配置
testConfig := &naming.Config{
    Source:  "etcd",
    Cluster: "test",
    // 其他配置...
}
```

### 健康检查

Nacos 和 etcd 实现都提供了健康检查机制：

- **Nacos**：通过心跳机制保持实例健康状态
- **etcd**：通过租约机制，当租约到期时实例自动被剔除

## 最佳实践

1. **服务命名规范**：使用一致的服务命名规范，建议采用域名风格（如 `com.example.service`）
2. **多实例部署**：为关键服务部署多个实例以提高可用性
3. **监听服务变更**：使用服务变更监听器实现动态路由和负载均衡
4. **优雅注销**：在服务优雅关闭时调用 `Deregister` 方法
5. **错误处理**：妥善处理注册、发现过程中的错误

## 故障处理

1. **连接失败**：在创建命名服务实例时会返回错误，需要妥善处理
2. **注册失败**：服务注册失败时应进行重试或告警
3. **服务发现失败**：实现服务降级或使用缓存的服务列表
4. **监听中断**：监听中断时会自动重连（特别是 etcd 实现）

## 依赖说明

- **etcd 实现**：
  - `go.etcd.io/etcd/client/v3`
  - `github.com/samber/lo`
  - `github.com/joshqu1985/lego/logs`

- **Nacos 实现**：
  - `github.com/nacos-group/nacos-sdk-go/v2`
  - `github.com/samber/lo`
  - `github.com/joshqu1985/lego/logs`