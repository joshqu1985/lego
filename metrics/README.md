# metrics
基于 github.com/prometheus/client_golang 封装的监控模块

## 使用方式
```go
import (
   "fmt"
   "net/http"

   "github.com/joshqu1985/lego/metrics"
)

var (
   serverMetricsTotal = metrics.NewCounter(metrics.CounterOpt{
      Namespace: "http_server",
      Name:      "code_total",
      Labels:    []string{"method", "code"},
   })
)

func main() {
   // 初始化监控 对外提供拉取接口 /metrics
   metrics.ServeHTTP()
   
   http.HandleFunc("/home", home)
   http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "Welcome!")
   serverMetricsTotal.Inc("GET:/", "200")
}
```
