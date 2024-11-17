package configor

import (
	"fmt"
	"testing"
	"time"

	"github.com/joshqu1985/lego/encoding"
)

type TestApolloData struct {
	C1 struct {
		Name string `toml:"name"`
		TTL  int    `toml:"ttl"`
	} `toml:"c1"`
	C2 struct {
		Enabled bool  `toml:"enabled"`
		Ports   []int `toml:"ports"`
	} `toml:"c2"`
}

func TestApolloConfig(t *testing.T) {
	var option options
	option.Encoding = encoding.New("toml")
	option.WatchChange = testWatcher
	config, err := NewApollo(&SourceConfig{
		Endpoints: []string{"http://127.0.0.1:8080/"},
		Cluster:   "default",
		AppId:     "helloworldid",
		Namespace: "application",
	}, option)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	data := &TestApolloData{}
	if err := config.Load(data); err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(data)

	time.Sleep(30 * time.Second)
}

func testWatcher(data ChangeSet) {
	fmt.Println("watch config change:\n", string(data.Value))
}
