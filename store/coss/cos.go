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

func NewCosClient(conf Config, option options) (*CosClient, error) {
	url, _ := url.Parse(conf.Endpoint)
	client := cos.NewClient(&cos.BaseURL{BucketURL: url}, &http.Client{
		Transport: &cos.AuthorizationTransport{SecretID: conf.AccessId, SecretKey: conf.AccessSecret},
	})

	return &CosClient{
		endpoint: conf.Endpoint,
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
	}, nil
}

func (this *CosClient) Get(ctx context.Context, key string, data io.Writer) error {
	resp, err := this.client.Object.Get(ctx, key, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)
	return err
}

func (this *CosClient) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := this.client.Object.Put(ctx, key, data, nil)
	if err != nil {
		return "", err
	}
	return this.client.Object.GetObjectURL(key).String(), nil
}

func (this *CosClient) GetFile(ctx context.Context, key, dstfile string) error {
	_, err := this.client.Object.GetToFile(ctx, key, dstfile, nil)
	return err
}

func (this *CosClient) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	_, err := this.client.Object.PutFromFile(ctx, key, srcfile, nil)
	if err != nil {
		return "", err
	}
	return this.client.Object.GetObjectURL(key).String(), nil
}

func (this *CosClient) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 1000,
		Marker:  nextToken,
	}

	nextMarker := ""
	resp, _, err := this.client.Bucket.Get(ctx, args)
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

func (this *CosClient) Head(ctx context.Context, key string) (ObjectMeta, error) {
	resp, err := this.client.Object.Head(ctx, key, nil)
	if err != nil {
		return ObjectMeta{}, err
	}

	return ObjectMeta{
		Key:           key,
		ContentLength: resp.ContentLength,
		ContentType:   resp.Header.Get("Content-Length"),
	}, nil
}
