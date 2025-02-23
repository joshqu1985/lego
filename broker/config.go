package broker

// Config 消息队列配置项
type Config struct {
	Source    string            `toml:"source" yaml:"source" json:"source"`
	Endpoints []string          `toml:"endpoints" yaml:"endpoints" json:"endpoints"`
	AppId     string            `toml:"app_id" yaml:"app_id" json:"app_id"`
	GroupId   string            `toml:"group_id" yaml:"group_id" json:"group_id"`
	Topics    map[string]string `toml:"topics" yaml:"topics" json:"topics"`
	AccessKey string            `toml:"access_key" yaml:"access_key" json:"access_key"`
	SecretKey string            `toml:"secret_key" yaml:"secret_key" json:"secret_key"`
}
