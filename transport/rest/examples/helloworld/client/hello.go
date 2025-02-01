package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rest"
)

func NewHelloworld() *Helloworld {
	target := "helloworld"
	c, err := rest.NewClient(target, rest.WithNaming(naming.Get()))
	if err != nil {
		log.Fatalf("rest.NewClient failed: %v", err)
		return nil
	}
	return &Helloworld{client: c}
}

type Helloworld struct {
	client *rest.Client
}

func (this *Helloworld) SayHello(ctx context.Context, data string) (string, error) {
	response, err := this.client.Request().
		SetQueryParams(map[string]string{
			"data": data,
		}).
		Get("/hello")
	if err != nil {
		log.Printf("could not request: %v", err)
		return "", err
	}
	if response.StatusCode() != 200 {
		return "", fmt.Errorf("invalid http code %d", response.StatusCode())
	}

	return string(response.Body()), nil
}
