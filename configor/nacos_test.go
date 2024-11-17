package configor

import (
	"testing"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

type TestNacosData struct {
	C1 struct {
		Name string `toml:"name"`
		Age  int    `toml:"age"`
	} `toml:"c1"`
	C2 struct {
		Enabled bool  `toml:"enabled"`
		Ports   []int `toml:"ports"`
	} `toml:"c2"`
}

func TestNacosConfig(t *testing.T) {
	var option options
	option.Encoding = encoding.New("toml")
	option.WatchChange = testWatcher
	config, err := NewNacos(&SourceConfig{
		Endpoints: []string{"127.0.0.1:8848"},
		Cluster:   "f5fc10a5-0f22-4188-ac9c-1e34023f6556",
		AppId:     "helloworldid",
	}, option)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	data := &TestNacosData{}
	if err := config.Load(data); err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(data)

	time.Sleep(30 * time.Second)
}
