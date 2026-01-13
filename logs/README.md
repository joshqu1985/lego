# 日志模块

logs 模块是一个基于 [go.uber.org/zap](https://github.com/uber-go/zap) 封装的高性能日志组件，提供了丰富的日志功能和灵活的配置选项，适用于各种 Go 语言应用程序的日志记录需求。

## 核心特性

- **高性能**：基于业界领先的 zap 日志库，提供极高的日志写入性能
- **多种输出目标**：支持控制台、文件和 Kafka 等多种日志输出方式
- **结构化日志**：支持结构化日志记录，便于日志分析和处理
- **上下文支持**：内置 trace-id 追踪支持，便于分布式系统问题定位
- **灵活配置**：提供丰富的配置选项，满足不同场景的需求
- **数据日志**：支持单独的数据流日志，与系统日志分离

## 支持的日志级别

- `DEBUG`：调试信息，通常在开发环境中使用
- `INFO`：一般信息，记录程序的正常运行状态
- `WARN`：警告信息，表示可能存在的问题但不影响程序继续运行
- `ERROR`：错误信息，表示程序出现了错误但仍可继续运行
- `PANIC`：严重错误，记录后会触发 panic
- `FATAL`：致命错误，记录后会导致程序退出

## 配置选项

### 环境设置

- `WithDevelopment()`：开发环境模式，使用更易读的控制台格式
- `WithProduction()`：生产环境模式，使用 JSON 格式记录日志

### 日志级别设置

- `WithLevel(level)`：设置日志级别，可选值为 `DEBUG`、`INFO`、`WARN`、`ERROR`、`PANIC`、`FATAL`

### 输出目标设置

- `WithWriterConsole()`：输出到控制台
- `WithWriterFile(fileOption)`：输出到文件，支持日志轮转

## 使用示例

### 基本初始化

```go
import (
    "github.com/joshqu1985/lego/logs"
)

func main() {
    // 开发环境，输出到控制台
    logs.Init(logs.WithDevelopment(), logs.WithWriterConsole())
    
    // 或者生产环境，输出到文件
    fileOption := &logs.FileOption{
        Filename:   "/var/log/myapp/server.log",
        MaxSize:    100,    // 单个文件最大尺寸（MB）
        MaxBackups: 10,     // 最大备份数量
        MaxAge:     28,     // 最大保留天数
        Compress:   true,   // 是否压缩
    }
    logs.Init(logs.WithProduction(), logs.WithWriterFile(fileOption))
}
```

### 多目标输出

```go
import (
    "github.com/joshqu1985/lego/logs"
)

func main() {
    // 同时输出到控制台、文件和 Kafka
    fileOption := &logs.FileOption{
        Filename: "/var/log/myapp/server.log",
        MaxSize: 100,
    }
    
    logs.Init(
        logs.WithProduction(),
        logs.WithWriterConsole(),
        logs.WithWriterFile(fileOption),
    )
}
```

### 日志输出 API

#### 基本日志函数

```go
// 简单文本日志
logs.Debug("调试信息")
logs.Info("一般信息")
logs.Warn("警告信息")
logs.Error("错误信息")
logs.Fatal("致命错误") // 记录后程序退出

// 格式化日志
logs.Debugf("调试信息: %s", value)
logs.Infof("一般信息: %d", count)
logs.Warnf("警告信息: %v", err)
logs.Errorf("错误信息: %v", err)
logs.Fatalf("致命错误: %v", err)
```

#### 带上下文的日志函数

这些函数会自动从 context 中提取 trace-id 并记录到日志中：

```go
import "context"

func handleRequest(ctx context.Context) {
    // 假设 ctx 中已经设置了 trace-id
    logs.Debugt(ctx, "处理请求", "path", "/api/users", "method", "GET")
    logs.Infot(ctx, "请求处理成功", "status", 200, "time_ms", 42)
    logs.Warnt(ctx, "请求参数异常", "param", "id", "value", "invalid")
    logs.Errort(ctx, "请求处理失败", "error", err)
    logs.Fatalt(ctx, "服务崩溃", "reason", "内存溢出")
}
```

#### 数据日志

用于记录业务数据日志，与系统日志分离：

```go
// 初始化数据日志
logs.InitData(logs.WithProduction(), logs.WithWriterFile(&logs.FileOption{
    Filename: "/var/log/myapp/data.log",
}))

// 输出数据日志
logs.Output("用户登录", "user_id", 12345, "timestamp", time.Now().Unix())
```

### Trace ID 支持

```go
import (
    "context"
    "github.com/joshqu1985/lego/logs"
)

func main() {
    // 创建带 trace-id 的 context
    ctx := context.WithValue(context.Background(), logs.CtxTraceKey("trace-id"), "123456789")
    
    // 使用带 trace-id 的日志函数
    logs.Infot(ctx, "处理请求")
    // 日志输出中会自动包含 TRACE-ID=123456789
    
    // 获取 context 中的 trace-id
    traceID := logs.GetTraceID(ctx)
    fmt.Println("当前 trace-id:", traceID)
}
```

## 最佳实践

1. **尽早初始化**：在应用程序启动时就初始化日志组件

2. **合理设置日志级别**：
   - 开发环境：使用 `DEBUG` 级别
   - 测试环境：使用 `INFO` 级别
   - 生产环境：使用 `WARN` 或 `ERROR` 级别

3. **结构化日志**：尽量使用键值对形式记录日志，便于后期分析

```go
// 推荐
logs.Infow("用户登录", "user_id", 123, "ip", "192.168.1.1", "success", true)

// 不推荐
logs.Infof("用户登录: user_id=%d, ip=%s, success=%v", 123, "192.168.1.1", true)
```

4. **错误处理**：在错误发生时记录详细信息，但避免在错误处理中再记录错误

5. **敏感信息保护**：不要在日志中记录密码、密钥等敏感信息

6. **日志文件管理**：
   - 设置合理的日志文件大小和保留时间
   - 启用日志压缩以节省磁盘空间
   - 定期清理过期日志

7. **性能考虑**：
   - 对于高频调用的日志，考虑使用条件判断避免不必要的日志处理
   - 大量日志输出时，考虑使用异步输出或批量处理

## 注意事项

1. 确保在使用日志函数前调用 `logs.Init()` 初始化日志组件

2. 生产环境中避免使用 `Debug` 级别日志，以减少性能开销

3. 使用 `WithWriterKafka` 时，确保传入的 producer 已经正确初始化

4. 日志文件路径需要确保应用程序有写权限

5. 避免在关键性能路径上记录过多日志，可能导致性能问题
