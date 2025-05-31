package coss

import (
	"context"
	"fmt"
	"io"
)

type Coss interface {
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

func New(conf Config, opts ...Option) (Coss, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	if option.BulkSize == 0 {
		option.BulkSize = int64(128 * 1024 * 1024) // 128MB
	}
	if option.Concurrency == 0 {
		option.Concurrency = 1
	}

	endpoint, ok := endpoints[conf.Source]
	if !ok {
		return nil, fmt.Errorf("unknown source: %s", conf.Source)
	}
	conf.Endpoint = "https://" + fmt.Sprintf(endpoint, conf.Region)

	var (
		coss Coss
		err  error
	)

	switch conf.Source {
	case "oss": // 阿里云
		coss, err = NewOssClient(conf, option)
	case "cos": // 腾讯云
		coss, err = NewCosClient(conf, option)
	case "obs": // 华为云
		coss, err = NewObsClient(conf, option)
	case "s3": // 亚马逊
		coss, err = NewS3Client(conf, option)
	default:
		coss, err = nil, fmt.Errorf("unknown source: %s", conf.Source)
	}
	return coss, err
}

type ObjectMeta struct {
	ContentType   string
	Key           string
	ContentLength int64
}
