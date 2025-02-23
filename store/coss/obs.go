package coss

import (
	"context"
	"io"

	"github.com/gabriel-vasile/mimetype"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type ObsClient struct {
	client   *obs.ObsClient
	option   options
	bucket   string
	endpoint string
}

func newObsClient(conf Config, option options) (*ObsClient, error) {
	client, err := obs.New(conf.AccessId, conf.AccessSecret, conf.Endpoint,
		obs.WithSignature(obs.SignatureObs))
	if err != nil {
		return nil, err
	}
	return &ObsClient{
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
		endpoint: conf.Endpoint,
	}, nil
}

func (this *ObsClient) Upload(ctx context.Context, key string, data io.Reader) (string, error) {
	mtype, err := mimetype.DetectReader(data)
	if err != nil {
		return "", err
	}

	args := &obs.PutObjectInput{}
	args.Bucket = this.bucket
	args.Key = key
	args.ContentType = mtype.String()
	args.Body = data

	_, err = this.client.PutObject(args)
	if err != nil {
		return "", err
	}
	return "https://" + this.bucket + "." + this.endpoint + "/" + key, nil
}

func (this *ObsClient) Download(ctx context.Context, key string, data io.Writer) error {
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

func (this *ObsClient) UploadFile(ctx context.Context, key, srcfile string) (string, error) {
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
	return "https://" + this.bucket + "." + this.endpoint + "/" + key, nil
}

func (this *ObsClient) DownloadFile(ctx context.Context, key, dstfile string) error {
	args := &obs.DownloadFileInput{}
	args.Bucket = this.bucket
	args.Key = key
	args.DownloadFile = dstfile
	args.PartSize = this.option.BulkSize
	args.TaskNum = this.option.Concurrency

	_, err := this.client.DownloadFile(args)
	return err
}
