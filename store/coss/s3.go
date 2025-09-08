package coss

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client   *s3.Client
	option   options
	endpoint string
	region   string
	bucket   string
	domain   string
}

func NewS3Client(conf Config, option options) (*S3Client, error) {
	creads := credentials.NewStaticCredentialsProvider(conf.AccessId, conf.AccessSecret, "")
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(creads))
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = conf.Region
		o.BaseEndpoint = aws.String(conf.Endpoint)
	})

	return &S3Client{
		endpoint: conf.Endpoint,
		client:   client,
		option:   option,
		region:   conf.Region,
		bucket:   conf.Bucket,
		domain:   conf.Domain,
	}, nil
}

func (this *S3Client) Get(ctx context.Context, key string, data io.Writer) error {
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

func (this *S3Client) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := this.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
		Body:   data,
	})
	if err != nil {
		return "", err
	}

	// https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-bucket-intro.html
	if this.domain != "" {
		return fmt.Sprintf("%s/%s", this.domain, key), nil
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", this.bucket, this.region, key), nil
}

func (this *S3Client) GetFile(ctx context.Context, key, dstfile string) error {
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

func (this *S3Client) PutFile(ctx context.Context, key, srcfile string) (string, error) {
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

	if this.domain != "" {
		return fmt.Sprintf("%s/%s", this.domain, key), nil
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", this.bucket, this.region, key), nil
}

func (this *S3Client) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &s3.ListObjectsV2Input{
		Bucket:  aws.String(this.bucket),
		Prefix:  aws.String(prefix),
		MaxKeys: aws.Int32(1000),
	}
	if nextToken != "" {
		args.ContinuationToken = aws.String(nextToken)
	}

	continueToken := ""
	resp, err := this.client.ListObjectsV2(ctx, args)
	if err != nil {
		return nil, "", err
	}
	if *resp.IsTruncated && resp.NextContinuationToken != nil {
		continueToken = *resp.NextContinuationToken
	}

	items := make([]ObjectMeta, 0, 1000)
	for _, content := range resp.Contents {
		items = append(items, ObjectMeta{Key: *content.Key, ContentLength: *content.Size})
	}
	return items, continueToken, nil
}

func (this *S3Client) Head(ctx context.Context, key string) (ObjectMeta, error) {
	resp, err := this.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return ObjectMeta{}, err
	}

	return ObjectMeta{
		Key:           key,
		ContentLength: *resp.ContentLength,
		ContentType:   *resp.ContentType,
	}, nil
}
