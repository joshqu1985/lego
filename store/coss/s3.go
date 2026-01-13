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

const (
	FORMAT_S3 = "https://%s." + ENDPOINT_S3 + "/%s"
)

type S3Client struct {
	client   *s3.Client
	endpoint string
	region   string
	bucket   string
	domain   string
	option   options
}

func NewS3Client(conf *Config, option options) (*S3Client, error) {
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

func (sc *S3Client) Get(ctx context.Context, key string, data io.Writer) error {
	resp, err := sc.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)

	return err
}

func (sc *S3Client) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := sc.client.PutObject(ctx, &s3.PutObjectInput{
		Body:   data,
		Key:    aws.String(key),
		Bucket: aws.String(sc.bucket),
	})
	if err != nil {
		return "", err
	}

	// https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-bucket-intro.html
	if sc.domain != "" {
		return fmt.Sprintf("%s/%s", sc.domain, key), nil
	}

	return fmt.Sprintf(FORMAT_S3, sc.bucket, sc.region, key), nil
}

func (sc *S3Client) GetFile(ctx context.Context, key, dstfile string) error {
	file, err := os.Create(dstfile)
	if err != nil {
		return err
	}
	defer file.Close()

	downloader := manager.NewDownloader(sc.client, func(d *manager.Downloader) {
		d.PartSize = sc.option.BulkSize
		d.Concurrency = sc.option.Concurrency
	})

	_, err = downloader.Download(ctx, file, &s3.GetObjectInput{
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
	})

	return err
}

func (sc *S3Client) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	reader, err := os.Open(srcfile)
	if err != nil {
		return "", err
	}

	uploader := manager.NewUploader(sc.client, func(u *manager.Uploader) {
		u.PartSize = sc.option.BulkSize
		u.Concurrency = sc.option.Concurrency
	})

	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(sc.bucket),
		Key:    aws.String(key),
		Body:   reader,
	})
	if err != nil {
		return "", err
	}

	if sc.domain != "" {
		return fmt.Sprintf("%s/%s", sc.domain, key), nil
	}

	return fmt.Sprintf(FORMAT_S3, sc.bucket, sc.region, key), nil
}

func (sc *S3Client) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &s3.ListObjectsV2Input{
		Bucket:  aws.String(sc.bucket),
		Prefix:  aws.String(prefix),
		MaxKeys: aws.Int32(1000),
	}
	if nextToken != "" {
		args.ContinuationToken = aws.String(nextToken)
	}

	continueToken := ""
	resp, err := sc.client.ListObjectsV2(ctx, args)
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

func (sc *S3Client) Head(ctx context.Context, key string) (ObjectMeta, error) {
	resp, err := sc.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(sc.bucket),
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
