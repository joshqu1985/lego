# logs
基于 go.uber.org/zap 封装的日志组件

## 使用
```go
import (
   "github.com/joshqu1985/lego/logs"
)

func main() {
   // 初始化日志 指定日志输出到控制台
   logs.Init(logs.WithDevelopment(), logs.WithWriterConsole())
   logs.Info("hello")
}
```
