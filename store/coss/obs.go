package coss

import (
	"context"
	"fmt"
	"io"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type ObsClient struct {
	client   *obs.ObsClient
	option   options
	endpoint string
	region   string
	bucket   string
}

func NewObsClient(conf Config, option options) (*ObsClient, error) {
	client, err := obs.New(conf.AccessId, conf.AccessSecret, conf.Endpoint,
		obs.WithSignature(obs.SignatureObs))
	if err != nil {
		return nil, err
	}

	return &ObsClient{
		endpoint: conf.Endpoint,
		region:   conf.Region,
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
	}, nil
}

func (this *ObsClient) Get(ctx context.Context, key string, data io.Writer) error {
	args := &obs.GetObjectInput{}
	args.Bucket = this.bucket
	args.Key = key

	resp, err := this.client.GetObject(args)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)
	return err
}

func (this *ObsClient) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	args := &obs.PutObjectInput{}
	args.Bucket = this.bucket
	args.Key = key
	args.Body = data

	_, err := this.client.PutObject(args)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.obs.%s.myhuaweicloud.com/%s", this.bucket, this.region, key), nil
}

func (this *ObsClient) GetFile(ctx context.Context, key, dstfile string) error {
	args := &obs.DownloadFileInput{}
	args.Bucket = this.bucket
	args.Key = key
	args.DownloadFile = dstfile
	args.PartSize = this.option.BulkSize
	args.TaskNum = this.option.Concurrency

	_, err := this.client.DownloadFile(args)
	return err
}

func (this *ObsClient) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	args := &obs.UploadFileInput{}
	args.Bucket = this.bucket
	args.Key = key
	args.UploadFile = srcfile
	args.PartSize = this.option.BulkSize
	args.TaskNum = this.option.Concurrency

	_, err := this.client.UploadFile(args)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.obs.%s.myhuaweicloud.com/%s", this.bucket, this.region, key), nil
}

func (this *ObsClient) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &obs.ListObjectsInput{}
	args.Bucket = this.bucket
	args.Prefix = prefix
	args.MaxKeys = 1000
	args.Marker = nextToken

	nextMarker := ""
	resp, err := this.client.ListObjects(args)
	if err != nil {
		return nil, "", err
	}
	if resp.IsTruncated {
		nextMarker = resp.NextMarker
	}

	items := make([]ObjectMeta, 0, 1000)
	for _, item := range resp.Contents {
		items = append(items, ObjectMeta{Key: item.Key, ContentLength: item.Size})
	}
	return items, nextMarker, nil
}

func (this *ObsClient) Head(ctx context.Context, key string) (ObjectMeta, error) {
	args := &obs.GetObjectMetadataInput{}
	args.Bucket = this.bucket
	args.Key = key

	resp, err := this.client.GetObjectMetadata(args)
	if err != nil {
		return ObjectMeta{}, err
	}

	return ObjectMeta{
		Key:           key,
		ContentLength: resp.ContentLength,
		ContentType:   resp.ContentType,
	}, nil
}
