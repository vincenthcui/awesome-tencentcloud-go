package tencentcloud

import (
	"context"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions"
)

type Interceptor func(ctx context.Context, action actions.Action, request, response interface{}, err error) error

func (c *Client) register(interceptor Interceptor) {
	c.interceptors = append(c.interceptors, interceptor)
}

func WithInterceptor(interceptor Interceptor) Option {
	return func(client *Client) {
		client.register(interceptor)
	}
}
