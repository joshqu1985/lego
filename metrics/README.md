# 指标监控模块

metrics 模块是一个基于 [github.com/prometheus/client_golang](https://github.com/prometheus/client_golang) 封装的监控指标收集库，提供了简洁易用的API，帮助开发者快速在应用中集成Prometheus监控。

## 核心特性

- **统一接口**：提供统一的指标创建和操作接口，隐藏底层实现细节
- **多种指标类型**：支持计数器（Counter）、仪表盘（Gauge）和直方图（Histogram）三种常用指标类型
- **灵活集成**：支持标准HTTP服务器和Gin框架两种集成方式
- **标签支持**：所有指标都支持标签（Labels），便于数据分类和筛选
- **开关控制**：提供全局开关控制，可动态启用或禁用指标收集
- **零开销设计**：当指标收集被禁用时，操作会被快速跳过，对性能影响极小

## 支持的指标类型

### 1. 计数器（Counter）

计数器是一种单调递增的指标，适用于记录请求数、错误数等只能增加不能减少的场景。

**特点**：
- 值只能增加或重置为零
- 适用于记录完成的请求数、发生的错误数等
- 支持自定义标签进行分类统计

**接口**：
- `Add(v float64, labels ...string)`：增加指定值
- `Inc(labels ...string)`：增加1

### 2. 仪表盘（Gauge）

仪表盘可以自由增减，适用于记录当前内存使用量、并发用户数等可以上下波动的场景。

**特点**：
- 值可以任意增加或减少
- 适用于记录当前状态，如内存使用量、活跃连接数等
- 支持自定义标签进行分类统计

**接口**：
- `Set(v float64, labels ...string)`：设置为指定值
- `Add(v float64, labels ...string)`：增加指定值
- `Sub(v float64, labels ...string)`：减少指定值

### 3. 直方图（Histogram）

直方图用于观察数据分布情况，适用于记录请求延迟、响应大小等需要统计分布的场景。

**特点**：
- 自动统计数据分布情况（分位数）
- 适用于记录延迟、大小等需要分析分布的数据
- 支持自定义分桶（buckets）和标签

**接口**：
- `Observe(v int64, labels ...string)`：观察一个整数值
- `ObserveFloat(v float64, labels ...string)`：观察一个浮点数值

## 集成方式

### 1. 标准HTTP服务器集成

```go
import (
    "net/http"
    "github.com/joshqu1985/lego/metrics"
)

func main() {
    // 初始化监控，对外提供 /metrics 接口
    metrics.ServeHTTP()
    
    // 启动HTTP服务器
    http.ListenAndServe(":8000", nil)
}
```

### 2. Gin框架集成

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/joshqu1985/lego/metrics"
)

func main() {
    router := gin.Default()
    
    // 将metrics集成到Gin路由
    metrics.ServeGIN(router)
    
    // 启动Gin服务器
    router.Run(":8000")
}
```

## 使用示例

### 创建和使用计数器

```go
import "github.com/joshqu1985/lego/metrics"

// 定义一个HTTP请求计数指标
var (
    httpRequestCounter = metrics.NewCounter(&metrics.CounterOpt{
        Namespace: "http_server",
        Name:      "request_total",
        Labels:    []string{"method", "path", "status"},
    })
)

func handleRequest(method, path string, statusCode int) {
    // 记录请求
    httpRequestCounter.Inc(method, path, string(statusCode))
    
    // 或者记录特定值
    httpRequestCounter.Add(2.5, method, path, string(statusCode))
}
```

### 创建和使用仪表盘

```go
import "github.com/joshqu1985/lego/metrics"

// 定义内存使用量指标
var (
    memoryUsageGauge = metrics.NewGauge(&metrics.GaugeOpt{
        Namespace: "app",
        Name:      "memory_usage_bytes",
        Labels:    []string{"type"},
    })
)

func updateMemoryMetrics() {
    // 设置内存使用量
    memoryUsageGauge.Set(1024*1024*50, "heap") // 50MB
    
    // 增加内存使用量
    memoryUsageGauge.Add(1024*1024*10, "heap") // 又增加了10MB
    
    // 减少内存使用量
    memoryUsageGauge.Sub(1024*1024*5, "heap") // 减少了5MB
}
```

### 创建和使用直方图

```go
import "github.com/joshqu1985/lego/metrics"

// 定义请求延迟指标
var (
    httpRequestDuration = metrics.NewHistogram(&metrics.HistogramOpt{
        Namespace: "http_server",
        Name:      "request_duration_seconds",
        Labels:    []string{"method", "path"},
        Buckets:   []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1.0, 5.0, 10.0},
    })
)

func measureRequestDuration(method, path string, durationMs int64) {
    // 观察延迟（毫秒转为秒）
    httpRequestDuration.ObserveFloat(float64(durationMs)/1000.0, method, path)
    
    // 或者直接观察整数值
    httpRequestDuration.Observe(durationMs, method, path)
}
```

### 完整的HTTP服务监控示例

```go
import (
    "fmt"
    "net/http"
    "time"
    "github.com/joshqu1985/lego/metrics"
)

var (
    // 请求计数
    requestCounter = metrics.NewCounter(&metrics.CounterOpt{
        Namespace: "http_server",
        Name:      "request_total",
        Labels:    []string{"method", "path", "status"},
    })
    
    // 请求延迟
    requestDuration = metrics.NewHistogram(&metrics.HistogramOpt{
        Namespace: "http_server",
        Name:      "request_duration_milliseconds",
        Labels:    []string{"method", "path"},
        Buckets:   []float64{1, 5, 10, 50, 100, 500, 1000, 5000},
    })
    
    // 活跃连接数
    activeConnections = metrics.NewGauge(&metrics.GaugeOpt{
        Namespace: "http_server",
        Name:      "active_connections",
        Labels:    []string{},
    })
)

func main() {
    // 初始化监控
    metrics.ServeHTTP()
    
    // 注册处理函数
    http.HandleFunc("/", withMetrics(rootHandler))
    http.HandleFunc("/api/users", withMetrics(usersHandler))
    
    fmt.Println("Server started on :8000")
    http.ListenAndServe(":8000", nil)
}

// 中间件：添加指标收集
func withMetrics(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 增加活跃连接数
        activeConnections.Add(1)
        defer activeConnections.Sub(1)
        
        // 记录开始时间
        startTime := time.Now()
        
        // 创建响应记录器来获取状态码
        rw := &responseWriter{ResponseWriter: w, statusCode: 200}
        
        // 处理请求
        handler(rw, r)
        
        // 计算延迟
        duration := time.Since(startTime).Milliseconds()
        
        // 记录指标
        requestCounter.Inc(r.Method, r.URL.Path, fmt.Sprintf("%d", rw.statusCode))
        requestDuration.Observe(duration, r.Method, r.URL.Path)
    }
}

// 用于捕获HTTP状态码的响应记录器
 type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (r *responseWriter) WriteHeader(statusCode int) {
    r.statusCode = statusCode
    r.ResponseWriter.WriteHeader(statusCode)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to the API!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Users endpoint")
}
```

## 最佳实践

1. **指标命名规范**：
   - 使用有意义的命名空间（Namespace）和名称（Name）
   - 名称应使用小写字母和下划线
   - 对于延迟指标，明确单位（如`_seconds`、`_milliseconds`）
   - 对于大小指标，明确单位（如`_bytes`、`_kilobytes`）

2. **标签使用**：
   - 只使用必要的标签，过多的标签会导致基数爆炸
   - 标签值应有限且可预测
   - 避免使用高基数的标签，如用户ID、请求ID等

3. **直方图分桶设计**：
   - 根据预期数据分布设计合适的分桶
   - 确保分桶覆盖可能的最小值和最大值
   - 对延迟指标，通常使用对数分布的分桶

4. **初始化时机**：
   - 在应用启动时初始化所有指标
   - 尽早调用`ServeHTTP()`或`ServeGIN()`启用指标端点

5. **性能考虑**：
   - 对于高频调用路径，考虑批量更新指标
   - 使用`Enabled()`函数在禁用状态时跳过不必要的处理
   - 避免在指标标签中使用复杂的字符串格式化操作

## 注意事项

1. 确保在使用指标前初始化监控（调用`ServeHTTP()`或`ServeGIN()`）

2. 指标一旦注册后，其配置（如标签数量和名称）不能更改

3. 避免在生产环境中使用过多的指标或标签，可能导致Prometheus性能下降

4. 对于直方图，自定义分桶时需要谨慎，不合理的分桶可能导致统计结果不准确

5. 当监控被禁用时（未调用初始化函数），所有指标操作仍然可以调用，但实际上不会记录任何数据
