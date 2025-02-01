package rest

import (
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joshqu1985/lego/transport/naming"
)

type Client struct {
	target string
	client *resty.Client
}

func NewClient(target string, opts ...Option) (*Client, error) {
	var option options
	for _, opt := range opts {
		opt(&option)
	}

	if option.Naming == nil {
		option.Naming = naming.NewPass(&naming.Config{})
	}
	if option.Timeout == 0 {
		option.Timeout = time.Duration(3) * time.Second
	}

	client := resty.New()
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		domain, err := ClientResolver(target, option.Naming)
		if err != nil {
			return err
		}
		req.URL = domain + req.URL

		u, err := url.Parse(req.URL)
		if err != nil {
			return err
		}

		method := req.Method + ":" + target + u.Path
		if err := ClientBreakerAllow(method); err != nil {
			return err
		}
		return nil
	})
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		u, err := url.Parse(resp.Request.URL)
		if err != nil {
			return err
		}

		method := resp.Request.Method + ":" + target + u.Path
		ClientBreakerMark(method, resp.StatusCode())
		ClientMetrics(method, resp.Time().Milliseconds(), resp.StatusCode())
		return nil
	})
	client.SetTimeout(option.Timeout)

	return &Client{
		target: target,
		client: client,
	}, nil
}

func (this *Client) Target() string {
	return this.target
}

func (this *Client) Request() *resty.Request {
	return this.client.R()
}
