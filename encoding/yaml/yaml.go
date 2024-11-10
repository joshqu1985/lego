package yaml

import (
	"io"
	"os"

	"github.com/ghodss/yaml"
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
	return yaml.Marshal(v)
}

func (this *Encoding) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func (this *Encoding) DecodeFile(file string, v any) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	data, err := io.ReadAll(fp)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}
