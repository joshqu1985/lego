package coss

import (
	"context"
	"os"
	"testing"
)

func InitS3() (*S3Client, error) {
	opts := options{}
	return NewS3Client(Config{
		Source:       "s3",
		Endpoint:     "http://127.0.0.1:9000",
		Region:       "cn",
		Bucket:       "test-bucket",
		AccessId:     "xx",
		AccessSecret: "xx",
	}, opts)
}

func TestS3List(t *testing.T) {
	s3, _ := InitS3()

	values, _, err := s3.List(context.Background(), "", "")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(values)
}

func TestS3Head(t *testing.T) {
	s3, _ := InitS3()

	meta, err := s3.Head(context.Background(), "1.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(meta)
}

func TestS3Get(t *testing.T) {
	s3, _ := InitS3()

	err := s3.Get(context.Background(), "1.txt", os.Stdout)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestS3Put(t *testing.T) {
	s3, _ := InitS3()

	file, err := os.Open("./test_image1.jpeg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer file.Close()

	addr, err := s3.Put(context.Background(), "test_imagex.jpeg", file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(addr)
}

func TestS3GetFile(t *testing.T) {
	s3, _ := InitS3()

	err := s3.GetFile(context.Background(), "07xkxkxkxk.png", "./tmp.jpg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestS3PutFile(t *testing.T) {
	s3, _ := InitS3()

	addr, err := s3.PutFile(context.Background(), "test_image1.jpeg", "./test_image1.jpeg")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(addr)
}
