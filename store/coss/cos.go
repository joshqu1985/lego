package coss

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosClient struct {
	client   *cos.Client
	option   options
	bucket   string
	endpoint string
}

func newCosClient(conf Config, option options) (*CosClient, error) {
	url, _ := url.Parse(conf.Endpoint)
	client := cos.NewClient(&cos.BaseURL{BucketURL: url}, &http.Client{
		Transport: &cos.AuthorizationTransport{SecretID: conf.AccessId, SecretKey: conf.AccessSecret},
	})
	return &CosClient{
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
		endpoint: conf.Endpoint,
	}, nil
}

func (this *CosClient) Upload(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := this.client.Object.Put(ctx, key, data, nil)
	if err != nil {
		return "", err
	}
	return this.client.Object.GetObjectURL(key).String(), nil
}

func (this *CosClient) Download(ctx context.Context, key string, data io.Writer) error {
	resp, err := this.client.Object.Get(ctx, key, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)
	return err
}

func (this *CosClient) UploadFile(ctx context.Context, key, srcfile string) (string, error) {
	_, err := this.client.Object.PutFromFile(ctx, key, srcfile, nil)
	if err != nil {
		return "", err
	}
	return this.client.Object.GetObjectURL(key).String(), nil
}

func (this *CosClient) DownloadFile(ctx context.Context, key, dstfile string) error {
	_, err := this.client.Object.GetToFile(ctx, key, dstfile, nil)
	return err
}
