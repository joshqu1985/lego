package coss

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gabriel-vasile/mimetype"
)

type S3Client struct {
	client   *s3.Client
	option   options
	endpoint string
	bucket   string
}

func newAwsClient(conf Config, option options) (*S3Client, error) {
	creads := credentials.NewStaticCredentialsProvider(conf.AccessId, conf.AccessSecret, "")
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(creads))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = conf.UsePathStyle
		o.Region = conf.Region
		o.BaseEndpoint = aws.String(conf.Endpoint)
	})

	return &S3Client{
		client:   client,
		option:   option,
		endpoint: conf.Endpoint,
		bucket:   conf.Bucket,
	}, nil
}

func (this *S3Client) Upload(ctx context.Context, key string, data io.Reader) (string, error) {
	mtype, err := mimetype.DetectReader(data)
	if err != nil {
		return "", err
	}

	_, err = this.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(this.bucket),
		Key:         aws.String(key),
		Body:        data,
		ContentType: aws.String(mtype.String()),
	})
	if err != nil {
		return "", err
	}
	return "https://" + this.bucket + "." + this.endpoint + "/" + key, nil
}

func (this *S3Client) Download(ctx context.Context, key string, data io.Writer) error {
	resp, err := this.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)
	return err
}

func (this *S3Client) UploadFile(ctx context.Context, key, srcfile string) (string, error) {
	reader, err := os.Open(srcfile)
	if err != nil {
		return "", err
	}

	uploader := manager.NewUploader(this.client, func(u *manager.Uploader) {
		u.PartSize = this.option.BulkSize
		u.Concurrency = this.option.Concurrency
	})

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}
	return "https://" + this.bucket + "." + this.endpoint + "/" + key, nil
}

func (this *S3Client) DownloadFile(ctx context.Context, key, dstfile string) error {
	file, err := os.Create(dstfile)
	if err != nil {
		return err
	}
	defer file.Close()

	downloader := manager.NewDownloader(this.client, func(d *manager.Downloader) {
		d.PartSize = this.option.BulkSize
		d.Concurrency = this.option.Concurrency
	})

	_, err = downloader.Download(ctx, file, &s3.GetObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (this *S3Client) headObject(key string) (*s3.HeadObjectOutput, error) {
	return this.client.HeadObject(context.Background(), &s3.HeadObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
	})
}
