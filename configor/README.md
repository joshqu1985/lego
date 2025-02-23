# configor
**配置中心** 对配置中心SDK做了封装, 提供统一接口, 方便快速使用和切换

## 使用方式
```golang
import (
   "github.com/joshqu1985/lego/configor"
)

func main() {
   config := configor.New("./config.toml",
      configor.WithToml(), config.WithWatch(watcher))
   if err := config.Load(&conf); err != nil {
      panic(err)
   }
}

func watcher(data configor.ChangeSet) {
   fmt.Println("watch config change:\n", string(data.Value))
}
```

## apollo
基于 github.com/apolloconfig/agollo/v4 封装
```toml
# config.toml
source = "apollo"
endpoints = ["http://127.0.0.1:8080/"]
cluster = "default"
app_id  = "helloworldid"
namespace = "application"
secret_key = ""
```

## etcd
基于 go.etcd.io/etcd/client/v3 封装
```toml
# config.toml
source = "etcd"
endpoints = ["http://127.0.0.1:9379"]
cluster = "default"
app_id  = "helloworldid"
access_key = ""
secret_key = ""
```

## nacos
基于 github.com/nacos-group/nacos-sdk-go/v2 封装
```toml
# config.toml
source = "nacos"
endpoints = ["http://127.0.0.1:8848"]
cluster = "f5fc10a5-0f22-4188-ac9c-1e34023f6556"
app_id  = "helloworldid"
access_key = ""
secret_key = ""
```

## local
```toml
# config.toml
[server]
  name = "helloworld"
  addr = ":8081"

[naming]
  source    = "etcd"
  endpoints = ["127.0.0.1:9379"]
```
