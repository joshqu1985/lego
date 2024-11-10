package json

import (
	"os"

	"github.com/bytedance/sonic"
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
	return sonic.Marshal(v)
}

func (this *Encoding) Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}

func (this *Encoding) DecodeFile(file string, v any) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	return sonic.ConfigDefault.NewDecoder(fp).Decode(v)
}
