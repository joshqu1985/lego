# REST 传输模块

## 概述

`transport/rest` 模块提供了基于 HTTP/REST 的服务通信框架，支持客户端和服务器端功能，集成了服务注册与发现、断路器、指标收集等功能，简化了微服务间的 HTTP 通信开发。

## 核心特性

- **统一的客户端和服务端接口**
- **服务注册与发现集成**
- **断路器保护机制**
- **性能指标收集**
- **优雅启动和关闭**
- **统一的错误处理和响应格式**
- **CORS 支持**
- **内置健康检查和调试功能**
- **函数式选项配置**

## 技术栈

- 客户端：基于 [go-resty/resty/v2](https://github.com/go-resty/resty) 封装
- 服务端：基于 [gin-gonic/gin](https://github.com/gin-gonic/gin) 框架
- 服务注册与发现：集成 `transport/naming` 模块
- 断路器：集成 `transport/breaker` 模块
- 指标收集：集成 `metrics` 模块

## 接口定义

### 客户端接口

```go
type Client struct {
    target string
    client *resty.Client
}

func NewClient(target string, opts ...Option) (*Client, error)
func (c *Client) Target() string
func (c *Client) Request() *resty.Request
```

### 服务端接口

```go
type Server struct {
    Name string
    Addr string
    // 内部字段...
}

type RouterRegister func(*Router)

type Option func(o *options)

func NewServer(name, addr string, opts ...Option) (*Server, error)
func (s *Server) Start() error
func (s *Server) Close() error
func (s *Server) AddMiddleware(inter gin.HandlerFunc)
```

### 响应结构

```go
type JSONResponse struct {
    Code int    `json:"code"`    // 状态码，格式为: 3位HTTP状态码+3位业务状态码
    Msg  string `json:"msg,omitempty"`  // 错误信息
    Data any    `json:"data,omitempty"` // 响应数据
}

type RouterHandler func(ctx *gin.Context) *JSONResponse
```

## 使用示例

### 创建和配置客户端

```go
import "github.com/joshqu1985/lego/transport/rest"

// 创建基本客户端
client, err := rest.NewClient("user-service")
if err != nil {
    log.Fatalf("Failed to create client: %v", err)
}

// 配置超时和服务发现
import "github.com/joshqu1985/lego/transport/naming"

namingClient, _ := naming.New(&naming.Config{
    Source:    "etcd",
    Endpoints: []string{"localhost:2379"},
})

client, err = rest.NewClient("user-service", 
    rest.WithTimeout(5*time.Second),
    rest.WithNaming(namingClient),
)

// 发送请求
resp, err := client.Request().
    SetPathParam("id", "123").
    SetQueryParam("active", "true").
    SetHeader("Content-Type", "application/json").
    Get("/users/{id}")

if err != nil {
    log.Printf("Request failed: %v", err)
    return
}

if resp.StatusCode() == 200 {
    var result map[string]interface{}
    if err := resp.Unmarshal(&result); err != nil {
        log.Printf("Failed to unmarshal response: %v", err)
    } else {
        log.Printf("Response: %v", result)
    }
}
```

### 创建和配置服务器

```go
import "github.com/joshqu1985/lego/transport/rest"

// 定义路由处理函数
func registerRoutes(r *rest.Router) {
    // 使用统一响应包装器
    r.GET("/users/:id", rest.ResponseWrapper(getUserHandler))
    r.POST("/users", rest.ResponseWrapper(createUserHandler))
}

// 路由处理函数示例
func getUserHandler(ctx *rest.Context) *rest.JSONResponse {
    id := ctx.Param("id")
    // 获取用户逻辑...
    if user == nil {
        return rest.ErrorResponse(404000, "User not found")
    }
    return rest.SuccessResponse(user)
}

// 创建服务器
server, err := rest.NewServer(
    "user-service",     // 服务名称
    "0.0.0.0:8080",     // 监听地址
    rest.WithRouters(registerRoutes),  // 注册路由
    rest.WithMetrics(),  // 启用指标收集
)

if err != nil {
    log.Fatalf("Failed to create server: %v", err)
}

// 启动服务器
if err := server.Start(); err != nil {
    log.Fatalf("Failed to start server: %v", err)
}
```

## 高级功能

### 服务注册与发现集成

REST 模块与 `transport/naming` 模块无缝集成，支持服务的自动注册和发现：

- **服务器端**：启动时自动注册服务，关闭时自动注销
- **客户端**：发送请求前自动从命名服务解析目标地址

```go
// 服务端注册示例
namingClient, _ := naming.New(&naming.Config{
    Source:    "nacos",
    Endpoints: []string{"localhost:8848"},
})

server, _ := rest.NewServer(
    "user-service",
    "0.0.0.0:8080",
    rest.WithNaming(namingClient),
    rest.WithRouters(registerRoutes),
)
```

### 断路器保护

客户端自动集成了断路器功能，当服务不可用时会自动熔断，防止级联故障：

- 基于 `transport/breaker` 模块实现
- 针对不同的 API 路径维护独立的断路器
- 500、503、504 等错误状态码会被视为失败
- 自动概率性尝试恢复

### 指标收集

模块提供了完善的指标收集功能：

**客户端指标**：
- `http_client_exec_duration`：请求耗时直方图
- `http_client_code_total`：请求状态码计数器

**服务端指标**：
- `http_server_exec_duration`：请求处理耗时直方图
- `http_server_code_total`：响应状态码计数器

通过 `WithMetrics()` 选项启用指标收集，指标可通过 `/metrics` 端点访问（Prometheus 格式）。

### 内置端点

REST 服务器自动注册了一系列有用的内置端点：

- **健康检查**：`/ping` - 返回 200 OK 和 "pong" 消息
- **性能分析**：`/debug/pprof/*` - Go 标准的性能分析端点
- **指标收集**：`/metrics` - 启用指标后可访问

## 中间件支持

### 内置中间件

- **CORS 支持**：处理跨域请求
- **恢复中间件**：捕获 panic 并返回 500 错误
- **指标中间件**：收集请求处理指标

### 自定义中间件

可以通过 `AddMiddleware` 方法添加自定义中间件：

```go
server.AddMiddleware(func(ctx *rest.Context) {
    // 自定义中间件逻辑
    ctx.Next()
})
```

## 统一响应格式

模块提供了统一的 JSON 响应格式和处理函数：

```go
// 成功响应
return rest.SuccessResponse(user)
// 返回: {"code": 200000, "data": {...}}

// 错误响应
return rest.ErrorResponse(400001, "Invalid parameter")
// 返回: {"code": 400001, "msg": "Invalid parameter"}
```

状态码格式为 6 位数字：前 3 位为 HTTP 状态码，后 3 位为业务状态码。

## 优雅启动和关闭

服务器支持优雅启动和关闭：

- 启动时自动注册服务
- 监听系统信号（SIGINT, SIGTERM）
- 接收到信号时优雅关闭（5秒超时）
- 关闭前注销服务，确保服务发现的一致性

```go
// 优雅关闭服务器
server.Close() // 向内部通道发送关闭信号
```

## 最佳实践

1. **服务命名规范**：使用一致的服务命名，便于服务发现
2. **统一错误处理**：使用 `ErrorResponse` 和 `SuccessResponse` 统一响应格式
3. **超时设置**：根据业务需要适当设置客户端超时
4. **指标监控**：启用并监控服务指标，及时发现性能问题
5. **错误重试**：对可重试的错误实现客户端重试逻辑
6. **日志记录**：记录关键操作的日志，便于问题排查

## 配置选项

| 选项 | 说明 | 默认值 |
|------|------|--------|
| WithTimeout | 设置客户端请求超时 | 3秒 |
| WithNaming | 设置服务注册与发现客户端 | 空实现 |
| WithRouters | 设置路由注册函数 | 必填 |
| WithMetrics | 启用指标收集 | false |

## 依赖说明

- **核心依赖**：
  - `github.com/go-resty/resty/v2` - HTTP 客户端
  - `github.com/gin-gonic/gin` - Web 框架
- **集成依赖**：
  - `github.com/joshqu1985/lego/transport/naming` - 服务注册与发现
  - `github.com/joshqu1985/lego/transport/breaker` - 断路器
  - `github.com/joshqu1985/lego/metrics` - 指标收集
