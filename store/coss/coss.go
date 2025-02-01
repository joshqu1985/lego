package coss

import (
	"fmt"
	"io"
)

type Coss interface {
	Upload(key string, data io.Reader) (string, error)
	Download(key string, data io.Writer) error
}

func New(conf Config, opts ...Option) (Coss, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if option.BulkSize == 0 {
		option.BulkSize = int64(10 * 1024 * 1024) // 10MB
	}

	var (
		coss Coss
		err  error
	)
	endpoint, ok := endpoints[conf.Source]
	if !ok {
		return nil, fmt.Errorf("unknown source: %s", conf.Source)
	}

	switch conf.Source {
	case "oss": // 阿里云
		conf.Endpoint = fmt.Sprintf(endpoint, conf.Region)
		coss, err = newAwsClient(conf, option)
	case "cos": // 腾讯云
		conf.Endpoint = fmt.Sprintf(endpoint, conf.Region)
		coss, err = newAwsClient(conf, option)
	case "obs": // 华为云
		conf.Endpoint = endpoint
		coss, err = newAwsClient(conf, option)
	default:
		coss, err = nil, fmt.Errorf("unknown source: %s", conf.Source)
	}
	return coss, err
}

type Config struct {
	Source       string `toml:"source" yaml:"source" json:"source"` // oss / cos / obs
	Region       string `toml:"region" yaml:"region" json:"region"`
	Bucket       string `toml:"bucket" yaml:"bucket" json:"bucket"`
	AccessKey    string `toml:"access_key" yaml:"access_key" json:"access_key"`
	AccessSecret string `toml:"access_secret" yaml:"access_secret" json:"access_secret"`

	Endpoint     string `toml:"-" yaml:"-" json:"-"`
	UsePathStyle bool   `toml:"-" yaml:"-" json:"-"`
}

var endpoints = map[string]string{
	"cos": "cos.%s.myqcloud.com",
	"oss": "oss-%s.aliyuncs.com",
	"obs": "obs.myhuaweicloud.com",
}
