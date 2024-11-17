package configor

import (
	"testing"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

type TestData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func TestLocalConfig(t *testing.T) {
	var option options
	option.Encoding = encoding.New("yaml")
	option.WatchChange = testWatcher
	config, err := NewLocal("./local_test.yaml", option)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	data := &TestData{}
	if err := config.Load(data); err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(data)

	time.Sleep(30 * time.Second)
}
