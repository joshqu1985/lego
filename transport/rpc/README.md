# RPC 传输模块

RPC模块是lego框架中的高性能远程过程调用组件，基于gRPC实现，提供了完整的客户端和服务器端功能，集成了服务注册发现、熔断器、指标监控等企业级特性。该模块旨在简化微服务之间的通信，提供稳定、可靠的RPC调用能力。

## 核心特性

- **统一接口**：提供简洁易用的客户端和服务器端API
- **多服务发现**：集成etcd、nacos等多种服务注册发现机制
- **熔断器保护**：内置断路器，防止系统级联故障
- **性能指标监控**：收集调用延迟、错误率等关键指标
- **超时控制**：支持配置请求超时时间
- **异常恢复**：自动捕获panic并转换为错误响应
- **优雅关闭**：支持服务优雅启动和关闭
- **健康检查**：内置健康检查端点
- **负载均衡**：客户端支持轮询负载均衡
- **自定义拦截器**：支持添加自定义的客户端和服务器端拦截器

## 技术栈

- **基础框架**：基于gRPC实现
- **服务发现**：集成etcd、nacos等
- **熔断器**：使用lego内置的breaker模块
- **监控**：使用lego内置的metrics模块

## 接口定义

### 客户端

`Client`结构体提供了gRPC客户端功能：

```go
type Client struct {
    target  string
    conn    *grpc.ClientConn
    builder gresolver.Builder
    // 其他内部字段
}
```

主要方法：
- `NewClient(target string, opts ...Option) (*Client, error)`: 创建新的客户端实例
- `Conn() *grpc.ClientConn`: 获取底层的gRPC连接
- `Target() string`: 获取目标服务

### 服务器端

`Server`结构体提供了gRPC服务器功能：

```go
type Server struct {
    Name string
    Addr string
    // 其他内部字段
}
```

主要方法：
- `NewServer(name, addr string, opts ...Option) (*Server, error)`: 创建新的服务器实例
- `Start() error`: 启动服务器
- `AddUnaryInterceptors(inter grpc.UnaryServerInterceptor)`: 添加一元拦截器
- `AddStreamInterceptors(inter grpc.StreamServerInterceptor)`: 添加流式拦截器

### 配置选项

`Option`函数允许配置客户端和服务器端：

- `WithTimeout(timeout time.Duration)`: 设置超时时间
- `WithNaming(n naming.Naming)`: 设置服务注册发现
- `WithRouters(f RouterRegister)`: 注册gRPC处理器
- `WithMetrics()`: 启用指标收集
```

## 拦截器

### 服务端拦截器

- **ServerUnaryRecover**：捕获并恢复 panic，确保服务稳定性
- **ServerStreamRecover**：流式 RPC 的 panic 恢复
- **ServerUnaryMetrics**：收集一元 RPC 的性能指标

### 客户端拦截器

- **ClientUnaryTimeout**：设置请求超时
- **ClientUnaryBreaker**：实现熔断器模式，防止服务雪崩
- **ClientUnaryMetrics**：收集客户端调用的性能指标

## 拦截器功能

### 客户端拦截器

- **超时拦截器**：控制请求超时时间
- **熔断器拦截器**：基于请求结果更新熔断器状态
- **指标拦截器**：收集调用延迟和错误码指标

### 服务器拦截器

- **异常恢复拦截器**：捕获panic并转换为错误响应
- **指标拦截器**：收集请求处理延迟和错误码指标

## 使用示例

### 服务端示例

```go
package main

import (
    "log"
    "google.golang.org/grpc"
    "github.com/joshqu1985/lego/transport/rpc"
    "github.com/joshqu1985/lego/transport/naming"
    pb "your/protobuf/package"
)

func main() {
    // 创建naming实例（示例使用etcd）
    naming, err := naming.New(&naming.Config{
        Source:    "etcd",
        Endpoints: []string{"localhost:2379"},
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建RPC服务器
    server, err := rpc.NewServer("service-name", "0.0.0.0:8080",
        rpc.WithNaming(naming),
        rpc.WithMetrics(),
        rpc.WithRouters(func(s *grpc.Server) {
            // 注册服务
            pb.RegisterYourServiceServer(s, &YourServiceImpl{})
        }))
    
    if err != nil {
        log.Fatal(err)
    }
    
    // 启动服务器
    if err := server.Start(); err != nil {
        log.Fatal(err)
    }
}

// YourServiceImpl 实现你的服务接口
type YourServiceImpl struct {
    pb.UnimplementedYourServiceServer
}

// 实现具体的RPC方法
func (s *YourServiceImpl) YourMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
    // 业务逻辑实现
    return &pb.YourResponse{Message: "Hello, " + req.Name}, nil
}
```

### 客户端示例

```go
package main

import (
    "context"
    "log"
    "time"
    "github.com/joshqu1985/lego/transport/rpc"
    "github.com/joshqu1985/lego/transport/naming"
    pb "your/protobuf/package"
)

func main() {
    // 创建naming实例（示例使用etcd）
    naming, err := naming.New(&naming.Config{
        Source:    "etcd",
        Endpoints: []string{"localhost:2379"},
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建RPC客户端
    client, err := rpc.NewClient("service-name", 
        rpc.WithNaming(naming),
        rpc.WithTimeout(3*time.Second))
    
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建gRPC客户端
    grpcClient := pb.NewYourServiceClient(client.Conn())
    
    // 调用服务
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    resp, err := grpcClient.YourMethod(ctx, &pb.YourRequest{Name: "World"})
    if err != nil {
        log.Fatal(err)
    }
    
    // 处理响应
    log.Printf("Response: %+v", resp)
}
```

## 健康检查

RPC 服务器自动启动 HTTP 健康检查服务，端口为 gRPC 端口 + 1。健康检查端点为 `/ping`，返回 `pong`。

```bash
curl http://localhost:8001/ping  # 假设 gRPC 服务运行在 8000 端口
```

## 指标收集

模块内置了指标收集功能，自动记录以下指标：

### 服务端指标

- `rpc_server_exec_duration`：RPC 调用耗时直方图
- `rpc_server_code_total`：RPC 调用结果计数器（按方法和状态码分类）

### 客户端指标

- `rpc_client_exec_duration`：客户端调用耗时直方图
- `rpc_client_code_total`：客户端调用结果计数器（按方法和状态码分类）

## 信号处理与优雅关闭

RPC模块提供了`Monitor`结构体用于处理系统信号，确保服务能够优雅关闭：

```go
type Monitor struct {
    callbacks []func()
    // 其他内部字段
}
```

主要方法：
- `NewMonitor() *Monitor`: 创建新的监控实例
- `AddShutdownCallback(fn func())`: 添加关闭回调函数

在服务器启动时，会自动创建一个`Monitor`实例，并注册关闭回调，确保在收到终止信号时能够注销服务并优雅关闭gRPC服务器。

## 高级功能

### 服务注册与发现集成

RPC模块与lego的naming模块无缝集成，支持etcd、nacos等多种服务注册发现机制。客户端可以自动发现服务实例，实现负载均衡和故障转移。

### 熔断器保护

内置的熔断器拦截器使用lego的breaker模块，自动监控请求失败率，在服务不可用时快速失败，防止系统过载。

### 性能指标

自动收集RPC调用的关键性能指标，包括调用延迟分布和错误码统计，方便进行性能监控和问题排查。

## 最佳实践

1. **合理配置超时时间**：根据实际业务场景设置适当的超时时间，避免长时间阻塞
2. **启用指标收集**：建议在生产环境中启用指标收集，便于监控服务性能
3. **使用服务注册发现**：在微服务架构中，优先使用etcd或nacos等服务注册发现机制
4. **合理处理错误**：正确处理和传递gRPC错误，便于问题排查
5. **添加自定义拦截器**：根据业务需求添加自定义的拦截器，如日志记录、链路追踪等
6. **实现优雅关闭**：确保在服务停止时能够释放资源，取消注册服务

## 故障处理

1. **连接失败**：检查目标服务是否正常运行，网络是否畅通
2. **服务发现失败**：确认服务注册发现组件配置正确，服务已注册
3. **熔断触发**：检查目标服务是否存在问题，调整熔断器参数
4. **超时错误**：检查网络延迟，考虑增加超时时间或优化服务性能

## 依赖说明

- **gRPC**：google.golang.org/grpc
- **服务发现**：依赖lego/transport/naming模块
- **熔断器**：依赖lego/transport/breaker模块
- **指标监控**：依赖lego/metrics模块

## 版本兼容性

- 支持Go 1.16+版本
- 与gRPC主流版本兼容
- 与lego框架其他模块协同工作
