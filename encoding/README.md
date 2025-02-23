# encoding
对json、yaml、xml数据序列化和反序列化封装，避免同一套业务代码出现不同版本甚至不同三方库的情况。

## json
基于 github.com/bytedance/sonic 封装
```golang
import (
   "github.com/joshqu1985/lego/encoding/json"
)

func main() {
   b, _ := json.Marshal(&data1)   

   _ = json.Unmarshal(b, &data2)

   _ = json.DecodeFile("./config.json", &data3)
}
```

## toml
基于 github.com/BurntSushi/toml 封装
```golang
import (
   "github.com/joshqu1985/lego/encoding/toml"
)

func main() {
   b, _ := toml.Marshal(&data1)   

   _ = toml.Unmarshal(b, &data2)

   _ = toml.DecodeFile("./config.toml", &data3)
}
```

## yaml
基于 github.com/ghodss/yaml 封装
```golang
import (
   "github.com/joshqu1985/lego/encoding/yaml"
)

func main() {
   b, _ := yaml.Marshal(&data1)   

   _ = yaml.Unmarshal(b, &data2)

   _ = yaml.DecodeFile("./config.yaml", &data3)
}
```
