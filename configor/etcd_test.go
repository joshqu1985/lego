package configor

import (
	"testing"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

type TestEtcdData struct {
	C1 struct {
		Name string `toml:"name"`
		Age  int    `toml:"age"`
	} `toml:"c1"`
	C2 struct {
		Enabled bool  `toml:"enabled"`
		Ports   []int `toml:"ports"`
	} `toml:"c2"`
}

func TestEtcdConfig(t *testing.T) {
	var option options
	option.Encoding = encoding.New("toml")
	option.WatchChange = testWatcher
	config, err := NewEtcd(&SourceConfig{
		Endpoints: []string{"http://127.0.0.1:9379"},
		Cluster:   "default",
		AppId:     "helloworldid",
	}, option)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	data := &TestEtcdData{}
	if err := config.Load(data); err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(data)

	time.Sleep(30 * time.Second)
}
