# 工具库模块

## 模块概述

utils模块是lego框架中的实用工具集合，提供了各种常用的功能组件，用于简化开发过程中常见的操作。这些工具模块设计为独立的、低耦合的组件，可以在框架内部或单独使用。

## 模块列表

| 模块名称 | 描述 | 文件路径 |
|---------|------|----------|
| crypto | 加密相关工具，提供MD5、HMAC、SHA1等加密算法 | utils/crypto/ |
| ip | IP地址相关工具，提供获取本地IP地址的功能 | utils/ip/ |
| pagetoken | 分页令牌工具，用于实现基于令牌的分页功能 | utils/pagetoken/ |
| routine | 协程工具，提供安全的协程创建和错误处理 | utils/routine/ |
| slice | 切片工具，提供通用的切片操作函数 | utils/slice/ |
| strings | 字符串工具，提供字符串和字节切片的零拷贝转换 | utils/strings/ |
| utime | 时间工具，提供扩展的时间处理功能 | utils/utime/ |

## 模块详细说明

### crypto 模块

提供常见的加密算法实现，包括MD5、HMAC和SHA1。

主要功能：
- `Md5(str string) (string, error)`: 计算字符串的MD5哈希值
- `Hmac(key, str string) (string, error)`: 使用MD5算法计算HMAC哈希值
- `Sha1(str string) (string, error)`: 计算字符串的SHA1哈希值

### ip 模块

提供获取本地IP地址的功能。

主要功能：
- `Local() string`: 获取本机的非回环IPv4地址

### pagetoken 模块

提供基于令牌的分页功能，使用Base64编码的JSON格式存储分页信息。

主要功能：
- `Encode(ptoken *PageToken) string`: 将分页信息编码为令牌字符串
- `Decode(s string, defaultSize int) (PageToken, error)`: 解码令牌字符串为分页信息

### routine 模块

提供安全的协程创建和错误处理功能，防止协程中的panic导致整个程序崩溃。

主要功能：
- `Go(fn func())`: 创建安全的协程，自动捕获panic
- `Safe(fn func()) (err error)`: 安全执行函数，捕获并返回panic
- `FunctionName(v any) string`: 获取函数的名称

### slice 模块

提供通用的切片操作函数，支持泛型。

主要功能：
- `Split[T any](list []T, size int) [][]T`: 将切片分割成指定大小的子切片
- `Contain[T ~string \| int \| int64](list []T, one T) bool`: 检查切片是否包含指定元素
- `Join[T ~string \| int \| int64](elems []T, sep string) string`: 将切片元素连接成字符串

### strings 模块

提供字符串和字节切片之间的零拷贝转换功能，适用于性能敏感的场景。

主要功能：
- `String2Bytes(s string) []byte`: 将字符串转换为字节切片（零拷贝）
- `Bytes2String(b []byte) string`: 将字节切片转换为字符串（零拷贝）

### utime 模块

提供扩展的时间处理功能，封装了Go标准库的time包，并添加了更多实用方法。

主要功能：
- 时间创建和格式化：`Now()`, `New(s string, layout ...string)`, `String(layout ...string)`
- 时间范围计算：`DayBegin()`, `DayEnd()`, `WeekBegin()`, `WeekEnd()`, `MonthBegin()`, `MonthEnd()`, `YearBegin()`, `YearEnd()`
- 时间判断：`InNumDays(n int)`
- 数据序列化支持：`MarshalJSON()`, `UnmarshalJSON()`, `Scan()`, `Value()`

## 使用示例

### crypto 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/crypto"
)

func main() {
    // 计算MD5哈希值
    md5Val, _ := crypto.Md5("hello world")
    fmt.Println("MD5:", md5Val)
    
    // 计算HMAC哈希值
    hmacVal, _ := crypto.Hmac("secret_key", "hello world")
    fmt.Println("HMAC:", hmacVal)
    
    // 计算SHA1哈希值
    sha1Val, _ := crypto.Sha1("hello world")
    fmt.Println("SHA1:", sha1Val)
}
```

### ip 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/ip"
)

func main() {
    // 获取本机IP
    localIP := ip.Local()
    fmt.Println("Local IP:", localIP)
}
```

### pagetoken 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/pagetoken"
)

func main() {
    // 创建分页令牌
    token := &pagetoken.PageToken{
        Mark: 100,
        Size: 20,
    }
    
    // 编码令牌
    encoded := pagetoken.Encode(token)
    fmt.Println("Encoded:", encoded)
    
    // 解码令牌
    decoded, _ := pagetoken.Decode(encoded, 10)
    fmt.Printf("Decoded: Mark=%d, Size=%d\n", decoded.Mark, decoded.Size)
}
```

### routine 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/routine"
)

func main() {
    // 安全创建协程
    routine.Go(func() {
        fmt.Println("Running in goroutine")
        // 即使这里发生panic，也不会导致程序崩溃
    })
    
    // 安全执行函数
    err := routine.Safe(func() {
        // 可能会panic的代码
    })
    if err != nil {
        fmt.Println("Error caught:", err)
    }
}
```

### slice 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/slice"
)

func main() {
    // 分割切片
    list := []int{1, 2, 3, 4, 5, 6, 7}
    chunks := slice.Split(list, 3)
    fmt.Println("Split chunks:", chunks)
    
    // 检查元素是否存在
    contains := slice.Contain([]string{"apple", "banana", "orange"}, "banana")
    fmt.Println("Contains banana:", contains)
    
    // 连接切片元素
    joined := slice.Join([]int{1, 2, 3, 4, 5}, ",")
    fmt.Println("Joined:", joined)
}
```

### strings 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/strings"
)

func main() {
    // 零拷贝转换字符串到字节切片
    s := "hello world"
    b := strings.String2Bytes(s)
    fmt.Println("String to bytes:", b)
    
    // 零拷贝转换字节切片到字符串
    newS := strings.Bytes2String(b)
    fmt.Println("Bytes to string:", newS)
}
```

### utime 模块示例

```go
package main

import (
    "fmt"
    "github.com/joshqu1985/lego/utils/utime"
)

func main() {
    // 创建当前时间
    now := utime.Now()
    fmt.Println("Now:", now.String())
    
    // 获取当天开始和结束时间
    dayBegin := now.DayBegin()
    dayEnd := now.DayEnd()
    fmt.Println("Day begin:", dayBegin.String())
    fmt.Println("Day end:", dayEnd.String())
    
    // 获取当月开始和结束时间
    monthBegin := now.MonthBegin()
    monthEnd := now.MonthEnd()
    fmt.Println("Month begin:", monthBegin.String())
    fmt.Println("Month end:", monthEnd.String())
    
    // 检查时间是否在指定天数内
    isRecent := now.InNumDays(7)
    fmt.Println("Is recent (7 days):", isRecent)
}
```

## 最佳实践

1. **安全使用协程**：在创建协程时，优先使用`routine.Go`或`routine.Safe`函数，避免未捕获的panic导致程序崩溃。

2. **性能优化**：在处理大型字符串或频繁进行字符串和字节切片转换的场景中，可以使用`strings`模块的零拷贝函数提高性能。

3. **分页实现**：对于需要分页的API，可以使用`pagetoken`模块实现基于令牌的分页，避免传统分页的性能问题。

4. **时间处理**：使用`utime`模块的扩展方法可以简化时间范围计算和格式化操作。

5. **切片操作**：对于需要频繁进行切片操作的代码，可以使用`slice`模块的泛型函数提高代码复用性。

6. **错误处理**：在使用`routine.Safe`函数时，应正确处理返回的错误，以便及时发现和解决问题。

## 注意事项

1. **零拷贝安全**：`strings`模块中的零拷贝转换函数在使用时需要注意，转换后的字节切片不应被修改，否则可能会导致未定义行为。

2. **时间解析**：`utime.New`函数在解析失败时会返回nil，使用时需要进行非空检查。

3. **协程安全**：虽然`routine`模块提供了panic保护，但仍需注意共享资源的并发访问安全。

4. **分页令牌安全**：`pagetoken`模块使用Base64编码的JSON，不提供加密保护，敏感信息不应直接存储在分页令牌中。

## 版本兼容性

- 支持Go 1.18+版本（由于使用了泛型）
- 与lego框架其他模块协同工作