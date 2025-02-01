package coss

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gabriel-vasile/mimetype"
)

type S3Client struct {
	client *s3.Client
	option options
	bucket string
}

func newAwsClient(conf Config, option options) (*S3Client, error) {
	creads := credentials.NewStaticCredentialsProvider(conf.AccessKey, conf.AccessSecret, "")
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
		client: client,
		option: option,
		bucket: conf.Bucket,
	}, nil
}

func (this *S3Client) Upload(key string, data io.Reader) (string, error) {
	uploader := manager.NewUploader(this.client, func(u *manager.Uploader) {
		u.PartSize = this.option.BulkSize
	})

	mtype, _ := mimetype.DetectReader(data)
	resp, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(this.bucket),
		Key:         aws.String(key),
		Body:        data,
		ContentType: aws.String(mtype.String()),
	})
	if err != nil {
		return "", err
	}
	return resp.Location, nil
}

func (this *S3Client) Download(key string, data io.Writer) error {
	head, err := this.headObject(key)
	if err != nil {
		return err
	}

	size := *head.ContentLength
	for beg := int64(0); beg < size; beg += this.option.BulkSize {
		end := beg + this.option.BulkSize - 1
		if end >= size {
			end = size - 1
		}

		if err := this.getObject(key, beg, end, data); err != nil {
			return err
		}
	}
	return nil
}

func (this *S3Client) headObject(key string) (*s3.HeadObjectOutput, error) {
	return this.client.HeadObject(context.Background(), &s3.HeadObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
	})
}

func (this *S3Client) getObject(key string, beg, end int64, data io.Writer) error {
	resp, err := this.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(this.bucket),
		Key:    aws.String(key),
		Range:  aws.String(fmt.Sprintf("bytes=%d-%d", beg, end)),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.CopyN(data, resp.Body, end-beg+1)
	return err
}
