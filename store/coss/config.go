package coss

type Config struct {
	Source       string `toml:"source" yaml:"source" json:"source"` // oss / cos / obs / s3
	Region       string `toml:"region" yaml:"region" json:"region"`
	Bucket       string `toml:"bucket" yaml:"bucket" json:"bucket"`
	AccessId     string `toml:"access_id" yaml:"access_id" json:"access_id"`
	AccessSecret string `toml:"access_secret" yaml:"access_secret" json:"access_secret"`
	Domain       string `toml:"domain" yaml:"domain" json:"domain"`

	Endpoint string `toml:"-" yaml:"-" json:"-"`
}

var endpoints = map[string]string{
	"cos": "%s.cos.%s.myqcloud.com",
	"oss": "oss-%s.aliyuncs.com",
	"obs": "obs.%s.myhuaweicloud.com",
	"s3":  "s3.%s.amazonaws.com",
}
