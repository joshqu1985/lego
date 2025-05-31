package coss

import (
	"context"
	"os"
	"testing"
)

func InitOss() (*OssClient, error) {
	opts := options{}
	return NewOssClient(Config{
		Source:       "oss",
		Endpoint:     "http://127.0.0.1:9000",
		Region:       "cn-beijing",
		Bucket:       "test-bucket",
		AccessId:     "xx",
		AccessSecret: "xx",
	}, opts)
}

func TestOssList(t *testing.T) {
	oss, _ := InitOss()

	prefix := "1"
	next := ""
	values, next, err := oss.List(context.Background(), prefix, next)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(values, next)
}

func TestOssHead(t *testing.T) {
	oss, _ := InitOss()

	meta, err := oss.Head(context.Background(), "1.png")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(meta)
}

func TestOssGet(t *testing.T) {
	oss, _ := InitOss()

	err := oss.Get(context.Background(), "1.png", os.Stdout)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestOssPut(t *testing.T) {
	oss, _ := InitOss()

	file, err := os.Open("./test_image1.jpeg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer file.Close()

	addr, err := oss.Put(context.Background(), "imagex.jpeg", file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(addr)
}

func TestOssGetFile(t *testing.T) {
	oss, _ := InitOss()

	err := oss.GetFile(context.Background(), "imagex.jpeg", "./tmpx.jpg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestOssPutFile(t *testing.T) {
	oss, _ := InitOss()

	addr, err := oss.PutFile(context.Background(), "imagey.jpeg", "./tmpx.jpg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(addr)
}
