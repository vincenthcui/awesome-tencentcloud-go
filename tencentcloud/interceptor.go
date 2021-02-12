package tencentcloud

import (
	"context"
)

type Interceptor func(ctx context.Context, action Action, request, response interface{}, err error) error

func (c *Client) register(interceptor Interceptor) {
	c.interceptors = append(c.interceptors, interceptor)
}

func WithInterceptor(interceptor Interceptor) Option {
	return func(client *Client) {
		client.register(interceptor)
	}
}
