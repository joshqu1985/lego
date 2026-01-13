package coss

import (
	"context"
	"fmt"
	"io"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

const (
	FORMAT_OSS = "https://%s." + ENDPOINT_OSS + "/%s"
)

type OssClient struct {
	client   *oss.Client
	endpoint string
	region   string
	bucket   string
	domain   string
	option   options
}

func NewOssClient(conf *Config, option options) (*OssClient, error) {
	provider := credentials.NewStaticCredentialsProvider(conf.AccessId, conf.AccessSecret, "")
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).
		WithRegion(conf.Region)
	client := oss.NewClient(cfg)

	return &OssClient{
		endpoint: conf.Endpoint,
		region:   conf.Region,
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
		domain:   conf.Domain,
	}, nil
}

func (osc *OssClient) Get(ctx context.Context, key string, data io.Writer) error {
	resp, err := osc.client.GetObject(ctx, &oss.GetObjectRequest{
		Bucket: oss.Ptr(osc.bucket),
		Key:    oss.Ptr(key),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)

	return err
}

func (osc *OssClient) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := osc.client.PutObject(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(osc.bucket),
		Key:    oss.Ptr(key),
		Body:   data,
	})
	if err != nil {
		return "", err
	}

	if osc.domain != "" {
		return fmt.Sprintf("%s/%s", osc.domain, key), nil
	}

	return fmt.Sprintf(FORMAT_OSS, osc.bucket, osc.region, key), nil
}

func (osc *OssClient) GetFile(ctx context.Context, key, dstfile string) error {
	downloader := osc.client.NewDownloader(func(d *oss.DownloaderOptions) {
		d.PartSize = osc.option.BulkSize
		d.ParallelNum = osc.option.Concurrency
	})
	opts := func(do *oss.DownloaderOptions) {}

	args := &oss.GetObjectRequest{
		Bucket: oss.Ptr(osc.bucket),
		Key:    oss.Ptr(key),
	}
	_, err := downloader.DownloadFile(ctx, args, dstfile, opts)

	return err
}

func (osc *OssClient) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	uploader := osc.client.NewUploader(func(u *oss.UploaderOptions) {
		u.PartSize = osc.option.BulkSize
		u.ParallelNum = osc.option.Concurrency
	})
	opts := func(u *oss.UploaderOptions) {}

	args := &oss.PutObjectRequest{
		Bucket: oss.Ptr(osc.bucket),
		Key:    oss.Ptr(key),
	}
	_, err := uploader.UploadFile(ctx, args, srcfile, opts)
	if err != nil {
		return "", err
	}

	if osc.domain != "" {
		return fmt.Sprintf("%s/%s", osc.domain, key), nil
	}

	return fmt.Sprintf(FORMAT_OSS, osc.bucket, osc.region, key), nil
}

func (osc *OssClient) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &oss.ListObjectsV2Request{
		Bucket:            oss.Ptr(osc.bucket),
		Prefix:            oss.Ptr(prefix),
		ContinuationToken: oss.Ptr(nextToken),
		MaxKeys:           5,
	}

	continueToken := ""
	resp, err := osc.client.ListObjectsV2(ctx, args)
	if err != nil {
		return nil, "", err
	}
	if resp.IsTruncated && resp.NextContinuationToken != nil {
		continueToken = *resp.NextContinuationToken
	}

	items := make([]ObjectMeta, 0, 1000)
	for _, content := range resp.Contents {
		items = append(items, ObjectMeta{Key: *content.Key, ContentLength: content.Size})
	}

	return items, continueToken, nil
}

func (osc *OssClient) Head(ctx context.Context, key string) (ObjectMeta, error) {
	resp, err := osc.client.HeadObject(ctx, &oss.HeadObjectRequest{
		Bucket: oss.Ptr(osc.bucket),
		Key:    oss.Ptr(key),
	})
	if err != nil {
		return ObjectMeta{}, err
	}

	return ObjectMeta{
		Key:           key,
		ContentLength: resp.ContentLength,
		ContentType:   *resp.ContentType,
	}, nil
}
