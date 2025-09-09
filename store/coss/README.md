# coss
**云对象存储服务** 对云对象存储服务SDK做了封装, 提供统一接口, 方便快速使用和切换

## 使用方式
```go
import (
   "github.com/joshqu1985/lego/store/coss"
)

func main() {
   // 初始化存储
   store, err := coss.New(conf)
   if err != nil {
      panic(err)
   }

   // 流式上传文件
   key := "test.txt"
   input, _ := os.Open("./test.txt")
   addr, _ := store.Put(context.Background(), "test.txt", input)

   // 流式下载文件
   key = "test.txt"
   output := os.Stdout
   _ = store.Get(context.Background(), key, output)

   // 上传文件
   key = "test.jpeg"
   src = "/tmp/test.jpeg"
   addr, _ := store.PutFile(context.Background(), key, src)

   // 下载文件
   key = "test.jpeg"
   dst := "/tmp/test.jpeg"
   _ = store.GetFile(context.Background(), key, dst)
}
```

## oss
阿里云对象存储服务 基于 github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss 封装
```golang
  // 配置信息
  conf := coss.Config{
    Source:       "oss",
    Region:       "cn-beijing",
    Bucket:       "example",
    AccessId:     "xx",
    AccessSecret: "xx",
    Domain:       "http://xxx.abc.com",
  }
```

## cos
腾讯云对象存储服务 基于 github.com/tencentyun/cos-go-sdk-v5 封装
```golang
  // 配置信息
  conf := coss.Config{
    Source:       "cos",
    Region:       "ap-beijing",
    Bucket:       "example-1234567",
    AccessId:     "xx",
    AccessSecret: "xx",
    Domain:       "http://xxx.abc.com",
  }
```

## obs
华为云对象存储服务 基于 github.com/huaweicloud/huaweicloud-sdk-go-obs/obs 封装
```golang
  // 配置信息
  conf := coss.Config{
    Source:       "obs",
    Region:       "cn-beijing",
    Bucket:       "example",
    AccessId:     "xx",
    AccessSecret: "xx",
    Domain:       "http://xxx.abc.com",
  }
```

## s3
亚马逊云对象存储服务 基于 github.com/aws/aws-sdk-go-v2/service/s3 封装
```golang
  // 配置信息
  conf := coss.Config{
    Source:       "s3",
    Region:       "cn",
    Bucket:       "example",
    AccessId:     "xx",
    AccessSecret: "xx",
    Domain:       "http://xxx.abc.com",
  }
```
