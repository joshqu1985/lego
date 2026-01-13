# 容器组件库

container 模块提供了一系列高性能的容器数据结构，用于满足不同场景下的数据存储和处理需求。这些容器实现了线程安全的操作，适用于高并发环境。

## 核心特性

- 线程安全：所有容器实现均支持并发安全操作
- 高性能：优化的内存使用和操作效率
- 易用性：简洁直观的 API 设计
- 功能丰富：涵盖常见的高级数据结构需求

## 容器类型

### 1. 批处理器 (Batcher)

批处理器用于将多个单个请求合并为批量请求，减少处理开销。适用于需要批量写入数据库、批量发送消息等场景。

#### 使用示例

```go
import (
    "github.com/joshqu1985/lego/container"
    "time"
    "encoding/json"
    "fmt"
)

func main() {
    // 500条 或者 1s 触发批量发送
    batcher := container.NewBatcher(
        container.WithBulkSize(500), 
        container.WithInterval(time.Second),
        container.WithQueueSize(1000),
    )
    go BatchSend(batcher.Queue())

    for i := 0; i < 1000; i++ {
        batcher.Put("hello")
    }
}

func BatchSend(queue chan []any) {
    for {
        vals := <-queue
        b, _ := json.Marshal(vals)
        fmt.Println("len: ", len(vals), " data: ", string(b))
    }
}
```

### 2. 位图 (BitMap)

位图是一种空间效率很高的数据结构，用于表示大量整数的存在与否。适用于去重、标记等场景。

#### 使用示例

```go
import (
    "github.com/joshqu1985/lego/container"
    "fmt"
)

func main() {
    // 创建一个能容纳1000个元素的位图
    bitmap := container.NewBitMap(1000)

    // 设置位
    _ = bitmap.Set(100)
    _ = bitmap.Set(200)

    // 检查位
    fmt.Println(bitmap.IsSet(100)) // 输出: true
    fmt.Println(bitmap.IsSet(150)) // 输出: false

    // 清除位
    _ = bitmap.Clear(100)
    fmt.Println(bitmap.IsSet(100)) // 输出: false
}
```

### 3. 布隆过滤器 (BloomFilter)

布隆过滤器是一种空间效率很高的概率性数据结构，用于判断一个元素是否在集合中。存在一定的误判率，但不会漏判。

#### 使用示例

```go
import (
    "github.com/joshqu1985/lego/container"
)

func main() {
    // 设置bit大小 10k
    bloom := container.NewBloomFilter(1024 * 10)
    
    // 添加
    bloom.Add([]byte("hello"))
    
    // 判断是否存在
    ok := bloom.Exist([]byte("world"))
}
```

### 4. 链表 (Linked)

线程安全的链表实现，适用于队列等场景。

#### 使用示例

```go
import (
    "github.com/joshqu1985/lego/container"
    "fmt"
)

func main() {
    // 创建链表
    linked := container.NewLinked()

    // 添加元素
    linked.Put("hello")
    linked.Put("world")

    // 获取元素
    data, err := linked.Get()
    if err == nil {
        fmt.Println(data) // 输出: hello
    }

    // 检查大小和是否为空
    fmt.Println(linked.Size())     // 输出: 1
    fmt.Println(linked.IsEmpty())  // 输出: false
}
```

### 5. 优先级队列 (Priority)

基于堆实现的优先级队列，支持按分数排序。分数越小的元素优先级越高。

#### 使用示例

```go
import (
    "github.com/joshqu1985/lego/container"
    "fmt"
)

func main() {
    // 创建优先级队列
    priority := container.NewPriority()

    // 添加带优先级的元素（分数越小优先级越高）
    priority.Put("task1", 3)
    priority.Put("task2", 1)
    priority.Put("task3", 2)

    // 获取顶部元素（不删除）
    top, score, _ := priority.Top()
    fmt.Println(top, score) // 输出: task2 1

    // 获取并删除最高优先级元素
    for !priority.IsEmpty() {
        task, _ := priority.Get()
        fmt.Println(task) // 依次输出: task2, task3, task1
    }
}
```

### 6. 滑动窗口 (SlidingWindow)

泛型滑动窗口实现，用于时间窗口内的数据统计和处理。

#### 使用示例

```go
import (
    "github.com/joshqu1985/lego/container"
    "time"
    "fmt"
)

// 定义一个简单的计数桶
type CountBucket struct {
    count int64
}

func (b *CountBucket) Add(cmd int64) {
    b.count += cmd
}

func (b *CountBucket) Reset() {
    b.count = 0
}

func main() {
    // 创建10个桶，每个桶代表100ms
    buckets := make([]CountBucket, 10)
    for i := range buckets {
        buckets[i] = CountBucket{}
    }

    // 创建滑动窗口
    window := container.NewSlidingWindow[CountBucket](100*time.Millisecond, buckets)

    // 添加数据
    for i := 0; i < 100; i++ {
        window.Add(1)
        time.Sleep(10 * time.Millisecond)
    }

    // 遍历窗口内的所有桶
    total := int64(0)
    window.Range(func(bucket CountBucket) {
        total += bucket.count
    })
    fmt.Println("Total in window:", total)
}
```

## 最佳实践

1. **并发安全**：所有容器都实现了线程安全的操作，可以直接在多个 goroutine 中使用

2. **内存效率**：
   - BitMap 和 BloomFilter 适用于需要高效存储大量布尔状态的场景
   - 对于频繁添加和获取操作，Linked 和 Priority 提供了良好的性能

3. **批量处理**：
   - Batcher 适用于需要将多个小操作合并为批量操作的场景，如数据库批量写入

4. **时间窗口统计**：
   - SlidingWindow 适用于限流、速率统计等基于时间窗口的场景

5. **布隆过滤器选择**：
   - 布隆过滤器的大小决定了容量和误判率
   - 对于假阳率要求高的场景，应选择较大的size

## 注意事项

1. BitMap 的索引不能超过初始化时设置的大小
2. BloomFilter 可能存在误判，但不会漏判
3. Batcher 的队列可能在高并发下阻塞，建议设置合理的队列大小
4. 滑动窗口的桶类型需要实现 IBucket 接口
