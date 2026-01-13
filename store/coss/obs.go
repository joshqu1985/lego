package coss

import (
	"context"
	"fmt"
	"io"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

const (
	FORMAT_OBS = "https://%s." + ENDPOINT_OBS + "/%s"
)

type ObsClient struct {
	client   *obs.ObsClient
	endpoint string
	region   string
	bucket   string
	domain   string
	option   options
}

func NewObsClient(conf *Config, option options) (*ObsClient, error) {
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
		domain:   conf.Domain,
	}, nil
}

func (obc *ObsClient) Get(ctx context.Context, key string, data io.Writer) error {
	args := &obs.GetObjectInput{}
	args.Bucket = obc.bucket
	args.Key = key

	resp, err := obc.client.GetObject(args)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)

	return err
}

func (obc *ObsClient) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	args := &obs.PutObjectInput{}
	args.Bucket = obc.bucket
	args.Key = key
	args.Body = data

	_, err := obc.client.PutObject(args)
	if err != nil {
		return "", err
	}

	if obc.domain != "" {
		return fmt.Sprintf("%s/%s", obc.domain, key), nil
	}

	return fmt.Sprintf(FORMAT_OBS, obc.bucket, obc.region, key), nil
}

func (obc *ObsClient) GetFile(ctx context.Context, key, dstfile string) error {
	args := &obs.DownloadFileInput{}
	args.Bucket = obc.bucket
	args.Key = key
	args.DownloadFile = dstfile
	args.PartSize = obc.option.BulkSize
	args.TaskNum = obc.option.Concurrency

	_, err := obc.client.DownloadFile(args)

	return err
}

func (obc *ObsClient) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	args := &obs.UploadFileInput{}
	args.Bucket = obc.bucket
	args.Key = key
	args.UploadFile = srcfile
	args.PartSize = obc.option.BulkSize
	args.TaskNum = obc.option.Concurrency

	_, err := obc.client.UploadFile(args)
	if err != nil {
		return "", err
	}

	if obc.domain != "" {
		return fmt.Sprintf("%s/%s", obc.domain, key), nil
	}

	return fmt.Sprintf(FORMAT_OBS, obc.bucket, obc.region, key), nil
}

func (obc *ObsClient) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &obs.ListObjectsInput{}
	args.Bucket = obc.bucket
	args.Prefix = prefix
	args.MaxKeys = 1000
	args.Marker = nextToken

	nextMarker := ""
	resp, err := obc.client.ListObjects(args)
	if err != nil {
		return nil, "", err
	}
	if resp.IsTruncated {
		nextMarker = resp.NextMarker
	}

	items := make([]ObjectMeta, 0, 1000)
	for i := range len(resp.Contents) {
		items = append(items, ObjectMeta{Key: resp.Contents[i].Key, ContentLength: resp.Contents[i].Size})
	}

	return items, nextMarker, nil
}

func (obc *ObsClient) Head(ctx context.Context, key string) (ObjectMeta, error) {
	args := &obs.GetObjectMetadataInput{}
	args.Bucket = obc.bucket
	args.Key = key

	resp, err := obc.client.GetObjectMetadata(args)
	if err != nil {
		return ObjectMeta{}, err
	}

	return ObjectMeta{
		Key:           key,
		ContentLength: resp.ContentLength,
		ContentType:   resp.ContentType,
	}, nil
}
