package coss

import (
	"context"
	"fmt"
	"io"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

type OssClient struct {
	client   *oss.Client
	option   options
	endpoint string
	region   string
	bucket   string
}

func NewOssClient(conf Config, option options) (*OssClient, error) {
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
	}, nil
}

func (this *OssClient) Get(ctx context.Context, key string, data io.Writer) error {
	resp, err := this.client.GetObject(ctx, &oss.GetObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)
	return err
}

func (this *OssClient) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := this.client.PutObject(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
		Body:   data,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.oss-%s.aliyuncs.com/%s", this.bucket, this.region, key), nil
}

func (this *OssClient) GetFile(ctx context.Context, key, dstfile string) error {
	downloader := this.client.NewDownloader(func(d *oss.DownloaderOptions) {
		d.PartSize = this.option.BulkSize
		d.ParallelNum = this.option.Concurrency
	})
	opts := func(do *oss.DownloaderOptions) {}

	args := &oss.GetObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
	}
	_, err := downloader.DownloadFile(ctx, args, dstfile, opts)
	return err
}

func (this *OssClient) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	uploader := this.client.NewUploader(func(u *oss.UploaderOptions) {
		u.PartSize = this.option.BulkSize
		u.ParallelNum = this.option.Concurrency
	})
	opts := func(u *oss.UploaderOptions) {}

	args := &oss.PutObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
	}
	_, err := uploader.UploadFile(ctx, args, srcfile, opts)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.oss-%s.aliyuncs.com/%s", this.bucket, this.region, key), nil
}

func (this *OssClient) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &oss.ListObjectsV2Request{
		Bucket:            oss.Ptr(this.bucket),
		Prefix:            oss.Ptr(prefix),
		ContinuationToken: oss.Ptr(nextToken),
		MaxKeys:           5,
	}

	continueToken := ""
	resp, err := this.client.ListObjectsV2(ctx, args)
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

func (this *OssClient) Head(ctx context.Context, key string) (ObjectMeta, error) {
	resp, err := this.client.HeadObject(ctx, &oss.HeadObjectRequest{
		Bucket: oss.Ptr(this.bucket),
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
