package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

var defaultEncoding = Encoding{}

func Marshal(v any) ([]byte, error) {
	return defaultEncoding.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return defaultEncoding.Unmarshal(data, v)
}

func DecodeFile(file string, v any) error {
	return defaultEncoding.DecodeFile(file, v)
}

type Encoding struct{}

func (this *Encoding) Marshal(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	defer buf.Reset()

	if err := toml.NewEncoder(buf).Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (this *Encoding) Unmarshal(data []byte, v any) error {
	_, err := toml.NewDecoder(bytes.NewReader(data)).Decode(v)
	return err
}

func (this *Encoding) DecodeFile(file string, v any) error {
	_, err := toml.DecodeFile(file, v)
	return err
}
