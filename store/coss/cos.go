package coss

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type CosClient struct {
	client   *cos.Client
	bucket   string
	endpoint string
	domain   string
	option   options
}

func NewCosClient(conf *Config, option options) (*CosClient, error) {
	url, _ := url.Parse(conf.Endpoint)
	client := cos.NewClient(&cos.BaseURL{BucketURL: url}, &http.Client{
		Transport: &cos.AuthorizationTransport{SecretID: conf.AccessId, SecretKey: conf.AccessSecret},
	})

	_, _, err := client.Service.Get(context.Background())
	if err != nil {
		return nil, err
	}

	return &CosClient{
		endpoint: conf.Endpoint,
		client:   client,
		option:   option,
		bucket:   conf.Bucket,
		domain:   conf.Domain,
	}, nil
}

func (cc *CosClient) Get(ctx context.Context, key string, data io.Writer) error {
	resp, err := cc.client.Object.Get(ctx, key, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(data, resp.Body)

	return err
}

func (cc *CosClient) Put(ctx context.Context, key string, data io.Reader) (string, error) {
	_, err := cc.client.Object.Put(ctx, key, data, nil)
	if err != nil {
		return "", err
	}

	if cc.domain != "" {
		return fmt.Sprintf("%s/%s", cc.domain, key), nil
	}

	return cc.client.Object.GetObjectURL(key).String(), nil
}

func (cc *CosClient) GetFile(ctx context.Context, key, dstfile string) error {
	_, err := cc.client.Object.GetToFile(ctx, key, dstfile, nil)

	return err
}

func (cc *CosClient) PutFile(ctx context.Context, key, srcfile string) (string, error) {
	_, err := cc.client.Object.PutFromFile(ctx, key, srcfile, nil)
	if err != nil {
		return "", err
	}

	if cc.domain != "" {
		return fmt.Sprintf("%s/%s", cc.domain, key), nil
	}

	return cc.client.Object.GetObjectURL(key).String(), nil
}

func (cc *CosClient) List(ctx context.Context, prefix, nextToken string) ([]ObjectMeta, string, error) {
	args := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 1000,
		Marker:  nextToken,
	}

	nextMarker := ""
	resp, _, err := cc.client.Bucket.Get(ctx, args)
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

func (cc *CosClient) Head(ctx context.Context, key string) (ObjectMeta, error) {
	resp, err := cc.client.Object.Head(ctx, key, nil)
	if err != nil {
		return ObjectMeta{}, err
	}

	return ObjectMeta{
		Key:           key,
		ContentLength: resp.ContentLength,
		ContentType:   resp.Header.Get("Content-Length"),
	}, nil
}
