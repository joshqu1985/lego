package coss

import (
	"context"
	"fmt"
	"io"
)

type Coss interface {
	Upload(ctx context.Context, key string, data io.Reader) (string, error)
	Download(ctx context.Context, key string, data io.Writer) error

	UploadFile(ctx context.Context, key, srcfile string) (string, error)
	DownloadFile(ctx context.Context, key, dstfile string) error
}

func New(conf Config, opts ...Option) (Coss, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if option.BulkSize == 0 {
		option.BulkSize = int64(256 * 1024 * 1024) // 256MB
	}
	if option.Concurrency == 0 {
		option.Concurrency = int64(3)
	}

	var (
		coss Coss
		err  error
	)
	endpoint, ok := endpoints[conf.Source]
	if !ok {
		return nil, fmt.Errorf("unknown source: %s", conf.Source)
	}
	conf.Endpoint = fmt.Sprintf(endpoint, conf.Region)

	switch conf.Source {
	case "oss": // 阿里云
		coss, err = newOssClient(conf, option)
	case "cos": // 腾讯云
		coss, err = newCosClient(conf, option)
	case "obs": // 华为云
		coss, err = newObsClient(conf, option)
	case "s3": // 亚马逊
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
	AccessId     string `toml:"access_id" yaml:"access_id" json:"access_id"`
	AccessSecret string `toml:"access_secret" yaml:"access_secret" json:"access_secret"`

	Endpoint     string `toml:"-" yaml:"-" json:"-"`
	UsePathStyle bool   `toml:"-" yaml:"-" json:"-"`
}

var endpoints = map[string]string{
	"cos": "cos.%s.myqcloud.com",
	"oss": "oss-%s.aliyuncs.com",
	"obs": "obs.%s.myhuaweicloud.com",
	"s3":  "s3-%s.amazonaws.com",
}
