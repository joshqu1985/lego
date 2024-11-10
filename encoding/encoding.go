package encoding

import (
	"lego/encoding/json"
	"lego/encoding/toml"
	"lego/encoding/yaml"
)

type Encoding interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
	DecodeFile(file string, v any) error
}

func New(enc string) Encoding {
	switch enc {
	case "json":
		return new(json.Encoding)
	case "yaml":
		return new(yaml.Encoding)
	case "toml":
		return new(toml.Encoding)
	}
	return nil
}
