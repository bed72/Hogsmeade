package clients

import "github.com/go-resty/resty/v2"

type RequestClient interface {
	Request() *resty.Request
}

type requestClient struct {
	key   string
	value string
}

func New(key, value string) RequestClient {
	return &requestClient{
		key:   key,
		value: value,
	}
}

func (r *requestClient) Request() *resty.Request {
	return resty.New().R().SetHeader(r.key, r.value)
}
