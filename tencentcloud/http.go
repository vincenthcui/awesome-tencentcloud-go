package tencentcloud

import (
	"net/http"
)

// http client interface, make it easy to mock
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func WithHttpClient(httpClient HttpClient) Option {
	return func(client *Client) {
		client.client = httpClient
	}
}
