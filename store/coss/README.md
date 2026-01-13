# 云对象存储服务 (COSS)

**云对象存储服务** 模块对主流云厂商的对象存储服务SDK进行了统一封装，提供简洁一致的接口，方便用户快速集成和切换不同的云存储服务。

## 核心特性

- **统一接口**：提供一致的API接口，屏蔽不同云厂商SDK的差异
- **多平台支持**：支持阿里云OSS、腾讯云COS、华为云OBS、AWS S3等主流云存储服务
- **灵活配置**：支持自定义存储路径、访问域名等配置
- **高级选项**：支持自定义块大小、并发数等上传下载参数
- **流式操作**：支持流式上传下载和文件操作
- **元数据管理**：支持查询对象元数据和列出对象

## 支持的云存储服务

- **OSS**: 阿里云对象存储服务
- **COS**: 腾讯云对象存储服务  
- **OBS**: 华为云对象存储服务
- **S3**: Amazon S3兼容的对象存储服务

## 统一接口定义

```go
type Coss interface {
    // Get 下载对象到io.Writer
    Get(ctx context.Context, key string, data io.Writer) error
    // Put 上传io.Reader到云存储，返回访问地址
    Put(ctx context.Context, key string, data io.Reader) (string, error)
    // GetFile 下载对象到本地文件
    GetFile(ctx context.Context, key, dstfile string) error
    // PutFile 上传本地文件到云存储，返回访问地址
    PutFile(ctx context.Context, key, srcfile string) (string, error)
    // List 列出桶内对象，支持前缀和分页
    List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error)
    // Head 查询对象元数据
    Head(ctx context.Context, key string) (ObjectMeta, error)
}

// ObjectMeta 对象元数据
type ObjectMeta struct {
    ContentType   string // 内容类型
    Key           string // 对象键
    ContentLength int64  // 内容长度
}
```

## 配置选项

### 基本配置

```go
type Config struct {
    Source       string // 云存储类型: oss/cos/obs/s3
    Region       string // 区域
    Bucket       string // 存储桶名称
    AccessId     string // 访问ID
    AccessSecret string // 访问密钥
    Domain       string // 自定义域名（可选）
}
```

### 高级选项

```go
// WithBulkSize 设置上传下载大文件时的块大小
func WithBulkSize(bulkSize int64) Option

// WithConcurrency 设置上传下载大文件时的并发数
func WithConcurrency(concurrency int) Option
```

## 使用示例

### 初始化

```go
import (
    "github.com/joshqu1985/lego/store/coss"
)

// 阿里云OSS配置
ossConfig := &coss.Config{
    Source:       "oss",
    Region:       "cn-beijing",
    Bucket:       "example-bucket",
    AccessId:     "your-access-id",
    AccessSecret: "your-access-secret",
    Domain:       "http://example.yourdomain.com", // 可选，自定义域名
}

// 初始化存储客户端
store, err := coss.New(ossConfig)
if err != nil {
    panic(err)
}

// 或使用高级选项
store, err = coss.New(ossConfig, 
    coss.WithBulkSize(64*1024*1024),  // 设置64MB的块大小
    coss.WithConcurrency(3),          // 设置3个并发
)
```

### 流式上传

```go
import (
    "context"
    "os"
)

// 打开本地文件作为输入流
input, err := os.Open("./example.txt")
if err != nil {
    panic(err)
}

def defer input.Close()

// 流式上传，返回访问地址
key := "documents/example.txt"
addr, err := store.Put(context.Background(), key, input)
if err != nil {
    panic(err)
}

fmt.Printf("上传成功，访问地址: %s\n", addr)
```

### 流式下载

```go
import (
    "context"
    "os"
)

// 创建文件作为输出流
output, err := os.Create("./download.txt")
if err != nil {
    panic(err)
}

def defer output.Close()

// 流式下载
key := "documents/example.txt"
err = store.Get(context.Background(), key, output)
if err != nil {
    panic(err)
}

fmt.Println("下载成功")
```

### 文件上传

```go
// 直接上传本地文件
key := "images/example.jpg"
srcFile := "/tmp/example.jpg"
addr, err := store.PutFile(context.Background(), key, srcFile)
if err != nil {
    panic(err)
}

fmt.Printf("文件上传成功，访问地址: %s\n", addr)
```

### 文件下载

```go
// 直接下载到本地文件
key := "images/example.jpg"
dstFile := "/tmp/download.jpg"
err = store.GetFile(context.Background(), key, dstFile)
if err != nil {
    panic(err)
}

fmt.Println("文件下载成功")
```

### 列出对象

```go
// 列出特定前缀的对象
prefix := "images/"
var nextToken string
objects, nextToken, err := store.List(context.Background(), prefix, nextToken)
if err != nil {
    panic(err)
}

for _, obj := range objects {
    fmt.Printf("文件: %s, 大小: %d, 类型: %s\n", 
        obj.Key, obj.ContentLength, obj.ContentType)
}

// 如果有更多数据，可以继续分页查询
if nextToken != "" {
    // 继续查询下一页
}
```

### 查询对象元数据

```go
// 查询对象元数据
key := "documents/example.txt"
meta, err := store.Head(context.Background(), key)
if err != nil {
    panic(err)
}

fmt.Printf("文件: %s\n", meta.Key)
fmt.Printf("大小: %d 字节\n", meta.ContentLength)
fmt.Printf("类型: %s\n", meta.ContentType)
```

## 各云存储服务配置示例

### 阿里云OSS

```go
conf := &coss.Config{
    Source:       "oss",
    Region:       "cn-beijing",       // 区域，如 cn-beijing, cn-shanghai
    Bucket:       "example-bucket",   // 存储桶名称
    AccessId:     "your-access-id",   // RAM访问ID
    AccessSecret: "your-access-secret", // RAM访问密钥
    Domain:       "http://example.yourdomain.com", // 可选，自定义域名
}
```

### 腾讯云COS

```go
conf := &coss.Config{
    Source:       "cos",
    Region:       "ap-beijing",       // 区域，如 ap-beijing, ap-guangzhou
    Bucket:       "example-bucket-123456789", // 存储桶名称
    AccessId:     "your-secret-id",   // 访问密钥ID
    AccessSecret: "your-secret-key",  // 访问密钥Key
    Domain:       "http://example.yourdomain.com", // 可选，自定义域名
}
```

### 华为云OBS

```go
conf := &coss.Config{
    Source:       "obs",
    Region:       "cn-north-4",       // 区域，如 cn-north-4, cn-east-3
    Bucket:       "example-bucket",   // 存储桶名称
    AccessId:     "your-access-key",  // 访问密钥ID
    AccessSecret: "your-secret-key",  // 访问密钥
    Domain:       "http://example.yourdomain.com", // 可选，自定义域名
}
```

### Amazon S3

```go
conf := &coss.Config{
    Source:       "s3",
    Region:       "us-east-1",        // 区域，如 us-east-1, eu-west-1
    Bucket:       "example-bucket",   // 存储桶名称
    AccessId:     "your-access-key-id", // 访问密钥ID
    AccessSecret: "your-secret-access-key", // 访问密钥
    Domain:       "http://example.yourdomain.com", // 可选，自定义域名
}
```

## 最佳实践

1. **错误处理**：始终检查API调用返回的错误，特别是网络操作
2. **上下文控制**：使用context控制请求超时和取消
3. **资源释放**：确保文件和流被正确关闭，避免资源泄漏
4. **并发配置**：对于大文件操作，可以适当调整WithBulkSize和WithConcurrency参数以提高性能
5. **密钥安全**：不要在代码中硬编码访问密钥，应从环境变量或配置管理系统获取
6. **路径设计**：使用有组织的键名结构（如日期/类型/文件）以便于管理和查询
7. **重试机制**：对于网络不稳定的环境，考虑添加重试逻辑

## 注意事项

1. 不同云存储服务的区域代码格式可能不同，请参考各云厂商的官方文档
2. 默认的块大小为128MB，并发数为1，可以根据实际需求调整
3. 自定义域名需要提前在各云厂商控制台进行配置
4. 使用第三方兼容S3协议的存储服务时，可能需要调整配置
5. 大文件操作时，建议使用WithBulkSize和WithConcurrency选项优化性能
6. 对于频繁访问的小文件，可以考虑使用CDN配合自定义域名来加速访问
