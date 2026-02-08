package httpx

import (
	"context"
	"net/http"
	"time"
)

type Client struct {
	HTTP *http.Client
}

func New() *Client {
	return &Client{HTTP: &http.Client{Timeout: 30 * time.Second}}
}

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
}

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (c *Client) Do(ctx context.Context, r Request) (Response, error) {
	req, err := http.NewRequestWithContext(ctx, r.Method, r.URL, nil)
	if err != nil {
		return Response{}, err
	}
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	res, err := c.HTTP.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	return Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       nil,
	}, nil
}
