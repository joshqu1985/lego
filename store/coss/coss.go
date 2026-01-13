package coss

import (
	"context"
	"fmt"
	"io"
)

const (
	DefaultBulkSize = 128 * 1024 * 1024 // 128MB

	COS = "cos"
	S3  = "s3"
	OSS = "oss"
	OBS = "obs"
)

type (
	Coss interface {
		// Get 下载
		Get(ctx context.Context, key string, data io.Writer) error
		// Put 上传
		Put(ctx context.Context, key string, data io.Reader) (string, error)
		// GetFile 下载到本地文件
		GetFile(ctx context.Context, key, dstfile string) error
		// PutFile 上传本地文件
		PutFile(ctx context.Context, key, srcfile string) (string, error)
		// List 列出桶内对象
		List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error)
		// Head 查询对象元数据
		Head(ctx context.Context, key string) (ObjectMeta, error)
	}

	ObjectMeta struct {
		ContentType   string
		Key           string
		ContentLength int64
	}
)

func New(conf *Config, opts ...Option) (Coss, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if option.BulkSize == 0 {
		option.BulkSize = int64(DefaultBulkSize)
	}
	if option.Concurrency == 0 {
		option.Concurrency = 1
	}

	endpoint, err := getCossEndpoints(conf)
	if err != nil {
		return nil, err
	}
	conf.Endpoint = endpoint

	switch conf.Source {
	case OSS: // 阿里云
		return NewOssClient(conf, option)
	case S3: // 亚马逊
		return NewS3Client(conf, option)
	case COS: // 腾讯云
		return NewCosClient(conf, option)
	case OBS: // 华为云
		return NewObsClient(conf, option)
	default:
		return nil, fmt.Errorf("unknown source: %s", conf.Source)
	}
}

func getCossEndpoints(conf *Config) (string, error) {
	endpoint, ok := endpoints[conf.Source]
	if !ok {
		return "", fmt.Errorf("unknown cloud source:%s", conf.Source)
	}

	if conf.Source == COS {
		return "https://" + fmt.Sprintf(endpoint, conf.Bucket, conf.Region), nil
	}

	return "https://" + fmt.Sprintf(endpoint, conf.Region), nil
}
