package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joshqu1985/lego/transport/naming"
	"github.com/joshqu1985/lego/transport/rest"
)

type Helloworld struct {
	client *rest.Client
}

func NewHelloworld(n naming.Naming) *Helloworld {
	target := "helloworld"
	c, err := rest.NewClient(target, rest.WithNaming(n))
	if err != nil {
		log.Printf("rest.NewClient failed: %v", err)

		return nil
	}

	return &Helloworld{client: c}
}

func (h *Helloworld) SayHello(ctx context.Context, data string) (string, error) {
	response, err := h.client.Request().
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
