package coss

import (
	"context"
	"io"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/gabriel-vasile/mimetype"
)

type OssClient struct {
	client   *oss.Client
	option   options
	bucket   string
	endpoint string
}

func newOssClient(conf Config, option options) (*OssClient, error) {
	provider := credentials.NewStaticCredentialsProvider(conf.AccessId, conf.AccessSecret, "")
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).
		WithRegion(conf.Region)
	client := oss.NewClient(cfg)

	return &OssClient{
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
		endpoint: conf.Endpoint,
	}, nil
}

func (this *OssClient) Upload(ctx context.Context, key string, data io.Reader) (string, error) {
	mtype, err := mimetype.DetectReader(data)
	if err != nil {
		return "", err
	}

	_, err = this.client.PutObject(ctx, &oss.PutObjectRequest{
		Bucket:      oss.Ptr(this.bucket),
		Key:         oss.Ptr(key),
		Body:        data,
		ContentType: oss.Ptr(mtype.String()),
	})
	if err != nil {
		return "", err
	}
	return "https://" + this.bucket + "." + this.endpoint + "/" + key, nil
}

func (this *OssClient) Download(ctx context.Context, key string, data io.Writer) error {
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

func (this *OssClient) UploadFile(ctx context.Context, key, srcfile string) (string, error) {
	uploader := this.client.NewUploader(func(u *oss.UploaderOptions) {
		u.PartSize = this.option.BulkSize
		u.ParallelNum = this.option.Concurrency
	})

	opts := func(u *oss.UploaderOptions) {
	}

	args := &oss.PutObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
	}
	_, err := uploader.UploadFile(ctx, args, srcfile, opts)
	if err != nil {
		return "", err
	}
	return "https://" + this.bucket + "." + this.endpoint + "/" + key, nil
}

func (this *OssClient) DownloadFile(ctx context.Context, key, dstfile string) error {
	downloader := this.client.NewDownloader(func(d *oss.DownloaderOptions) {
		d.PartSize = this.option.BulkSize
		d.ParallelNum = this.option.Concurrency
	})

	opts := func(do *oss.DownloaderOptions) {
	}

	args := &oss.GetObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
	}

	_, err := downloader.DownloadFile(ctx, args, dstfile, opts)
	return err
}

func (this *OssClient) headObject(key string) (*oss.HeadObjectResult, error) {
	return this.client.HeadObject(context.Background(), &oss.HeadObjectRequest{
		Bucket: oss.Ptr(this.bucket),
		Key:    oss.Ptr(key),
	})
}
