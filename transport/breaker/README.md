# 断路器模块

## 概述

`transport/breaker` 模块实现了基于 SRE（Site Reliability Engineering）算法的断路器模式，用于在分布式系统中防止级联故障，提高系统稳定性。断路器可以监控服务调用的成功率，并在失败率过高时自动熔断，避免持续向不可用的服务发送请求。

## 核心特性

- 统一的断路器接口定义
- 基于 SRE 算法的断路器实现
- 支持滑动窗口统计请求成功率
- 自动概率性探测恢复机制
- 线程安全的全局断路器管理

## 接口定义

### Breaker 接口

```go
type Breaker interface {
    Allow() bool      // 判断是否允许请求通过
    MarkPass()        // 标记请求成功
    MarkFail()        // 标记请求失败
}
```

## 实现原理

### SRE 断路器算法

SRE 断路器基于以下核心公式：

```
requests = k * access
```

其中：
- `k` 是倍率系数（默认为 1.8）
- `access` 是成功请求数
- `requests` 是预期的总请求数阈值

当实际请求数超过阈值时，断路器会根据失败率计算一个拒绝概率，概率性地拒绝部分请求，同时允许少量请求通过以探测服务是否恢复。

### 滑动窗口统计

使用 `container.SlidingWindow` 实现请求统计，默认配置：
- 窗口大小：20 个桶
- 时间窗口：300ms
- 最小请求数：100（低于此值时断路器不会触发熔断）

## 使用示例

### 基本用法

```go
import "github.com/joshqu1985/lego/transport/breaker"

// 获取或创建名为 "my_service" 的断路器
b := breaker.Get("my_service")

// 在调用服务前检查是否允许请求通过
if !b.Allow() {
    // 断路器已触发，拒绝请求
    return errors.New("service unavailable")
}

// 调用服务
err := callService()
if err != nil {
    // 标记请求失败
    b.MarkFail()
    return err
}

// 标记请求成功
b.MarkPass()
return nil
```

### 在 RPC 客户端中使用

```go
func (c *Client) Call(ctx context.Context, method string, req, resp interface{}) error {
    // 使用方法名作为断路器名称
    b := breaker.Get(method)
    
    if !b.Allow() {
        return errors.New("service circuit breaker triggered")
    }
    
    err := c.client.Call(ctx, method, req, resp)
    if err != nil {
        b.MarkFail()
        return err
    }
    
    b.MarkPass()
    return nil
}
```

## 高级配置

当前 SRE 断路器实现的默认参数如下：

| 参数 | 说明 | 默认值 |
|------|------|--------|
| k | 倍率系数 | 1.8 |
| minRequest | 最小请求数阈值 | 100 |
| slidingWindow | 滑动窗口大小 | 20 桶，300ms |

要自定义这些参数，需要修改代码中的 `newSREBreaker` 函数。

## 最佳实践

1. **为不同服务使用不同的断路器**：使用服务名或方法名作为断路器的唯一标识
2. **合理设置请求标记**：正确调用 `MarkPass` 和 `MarkFail` 方法
3. **处理熔断情况**：在断路器触发时提供合理的降级策略
4. **监控断路器状态**：建议添加监控以跟踪断路器的触发情况

## 性能考虑

- 断路器操作是线程安全的
- 统计数据存储在内存中，访问高效
- 内部使用读写锁减少竞争
- 支持高并发场景下的使用

## 依赖说明

- `github.com/joshqu1985/lego/container`：提供滑动窗口实现
- `golang.org/x/exp/rand`：提供随机数生成功能