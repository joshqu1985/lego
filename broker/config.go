package broker

// Config 消息队列配置项.
type Config struct {
	Topics    map[string]string `json:"topics"     toml:"topics"     yaml:"topics"`
	Source    string            `json:"source"     toml:"source"     yaml:"source"`
	AppId     string            `json:"app_id"     toml:"app_id"     yaml:"app_id"`
	GroupId   string            `json:"group_id"   toml:"group_id"   yaml:"group_id"`
	AccessKey string            `json:"access_key" toml:"access_key" yaml:"access_key"`
	SecretKey string            `json:"secret_key" toml:"secret_key" yaml:"secret_key"`
	Endpoints []string          `json:"endpoints"  toml:"endpoints"  yaml:"endpoints"`
}
