# 编码模块

encoding 模块提供了统一的数据序列化和反序列化接口，支持多种主流的数据格式（JSON、YAML、TOML）。该模块通过封装高性能的第三方库，为应用程序提供一致的编码解码体验，避免在同一套业务代码中使用不同版本或不同库的情况。

## 核心特性

- **统一接口**：所有编码格式实现相同的接口，便于在不同格式间切换
- **高性能**：使用业界领先的第三方库实现，如 JSON 采用 ByteDance 的 sonic 库
- **简洁API**：提供直观易用的 API，支持内存数据和文件数据的编解码
- **一致性**：确保在不同编码格式间保持一致的行为和错误处理

## 接口定义

### 统一编码接口

```go
type Encoding interface {
    // 将对象序列化为字节数组
    Marshal(v any) ([]byte, error)
    // 将字节数组反序列化为对象
    Unmarshal(data []byte, v any) error
    // 从文件中解码数据到对象
    DecodeFile(file string, v any) error
}
```

### 工厂函数

```go
// 根据编码名称创建对应的编码实例
func New(enc string) Encoding
```

## 支持的编码格式

### 1. JSON

基于 [github.com/bytedance/sonic](https://github.com/bytedance/sonic) 封装，这是一个高性能的 JSON 库。

**特点**：
- 极高的序列化和反序列化性能
- 完全兼容标准库 json 包的 API
- 支持流式处理大型 JSON 数据

**使用示例**：

```go
import (
    "github.com/joshqu1985/lego/encoding/json"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    // 序列化
    person := Person{Name: "张三", Age: 30, Email: "zhangsan@example.com"}
    b, err := json.Marshal(&person)
    if err != nil {
        // 处理错误
    }
    
    // 反序列化
    var decodedPerson Person
    err = json.Unmarshal(b, &decodedPerson)
    if err != nil {
        // 处理错误
    }
    
    // 从文件中解码
    var config Person
    err = json.DecodeFile("./config.json", &config)
    if err != nil {
        // 处理错误
    }
}
```

### 2. YAML

基于 [github.com/ghodss/yaml](https://github.com/ghodss/yaml) 封装，支持 YAML 1.2 标准。

**特点**：
- 支持丰富的 YAML 语法特性
- 可读性高，适合配置文件
- 支持注释和复杂的数据结构

**使用示例**：

```go
import (
    "github.com/joshqu1985/lego/encoding/yaml"
)

type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        URL      string `yaml:"url"`
        Username string `yaml:"username"`
        Password string `yaml:"password"`
    } `yaml:"database"`
}

func main() {
    // 序列化
    config := Config{
        // 初始化配置...
    }
    b, err := yaml.Marshal(&config)
    if err != nil {
        // 处理错误
    }
    
    // 反序列化
    var decodedConfig Config
    err = yaml.Unmarshal(b, &decodedConfig)
    if err != nil {
        // 处理错误
    }
    
    // 从文件中解码
    var fileConfig Config
    err = yaml.DecodeFile("./config.yaml", &fileConfig)
    if err != nil {
        // 处理错误
    }
}
```

### 3. TOML

基于 [github.com/BurntSushi/toml](https://github.com/BurntSushi/toml) 封装，支持 TOML 格式。

**特点**：
- 专为配置文件设计，语法简洁明了
- 支持表、数组、注释等特性
- 严格的类型定义，避免配置歧义

**使用示例**：

```go
import (
    "github.com/joshqu1985/lego/encoding/toml"
)

type AppConfig struct {
    Title   string
    Owner   OwnerInfo
    Database DatabaseConfig
    Servers map[string]Server
}

type OwnerInfo struct {
    Name string
    Org  string
    Bio  string
}

type DatabaseConfig struct {
    Server  string
    Ports   []int
    Enabled bool
}

type Server struct {
    IP string
    DC string
}

func main() {
    // 序列化
    config := AppConfig{
        // 初始化配置...
    }
    b, err := toml.Marshal(&config)
    if err != nil {
        // 处理错误
    }
    
    // 反序列化
    var decodedConfig AppConfig
    err = toml.Unmarshal(b, &decodedConfig)
    if err != nil {
        // 处理错误
    }
    
    // 从文件中解码
    var fileConfig AppConfig
    err = toml.DecodeFile("./config.toml", &fileConfig)
    if err != nil {
        // 处理错误
    }
}
```

## 使用统一接口

encoding 模块提供了一个统一的工厂函数，可以根据格式名称动态创建对应的编码实例，适用于需要根据配置动态选择编码格式的场景。

```go
import (
    "github.com/joshqu1985/lego/encoding"
)

func main() {
    // 根据配置选择编码格式
    encType := "json" // 可以从配置中读取
    encoder := encoding.New(encType)
    
    if encoder == nil {
        // 处理不支持的编码格式
    }
    
    // 使用统一接口进行编解码
    data := map[string]interface{}{
        "key": "value",
        "num": 42,
    }
    
    // 序列化
    b, err := encoder.Marshal(data)
    if err != nil {
        // 处理错误
    }
    
    // 反序列化
    var decodedData map[string]interface{}
    err = encoder.Unmarshal(b, &decodedData)
    if err != nil {
        // 处理错误
    }
    
    // 从文件中解码
    err = encoder.DecodeFile("./config."+encType, &decodedData)
    if err != nil {
        // 处理错误
    }
}
```

## 最佳实践

1. **选择合适的编码格式**：
   - JSON：网络传输、API响应、简单配置
   - YAML：复杂配置文件、需要注释和可读性高的场景
   - TOML：配置文件，尤其是需要严格类型定义的场景

2. **统一编码库**：在整个项目中统一使用 encoding 模块，避免引入多个版本或不同的编码库

3. **错误处理**：总是检查编解码操作返回的错误，特别是在处理外部数据时

4. **大型数据处理**：对于大型数据文件，可以考虑直接使用 DecodeFile 方法，避免将整个文件加载到内存中

5. **结构体标签**：使用适当的结构体标签（如 json:"name"、yaml:"name"、toml:"name"）控制序列化行为

## 注意事项

1. 不同编码格式对数据类型的支持有所差异，某些在 JSON 中支持的数据类型可能在其他格式中有不同的表示

2. 对于同一结构体进行不同格式的序列化和反序列化时，确保结构体标签与目标格式匹配

3. 处理敏感数据时，注意编码后的内容可能包含明文敏感信息

4. 对于特别大的数据集，考虑使用流式处理以减少内存占用
