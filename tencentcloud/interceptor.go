package tencentcloud

import (
	"context"
)

// Interceptor hooks after request, like a pipeline
// OnNetworkFailure and OnRateLimitExceeded are default interceptors
type Interceptor func(ctx context.Context, action Action, request, response interface{}, err error) error

func (c *Client) register(interceptor Interceptor) {
	c.interceptors = append(c.interceptors, interceptor)
}

// Option to append interceptor
func WithInterceptor(interceptor Interceptor) Option {
	return func(client *Client) {
		client.register(interceptor)
	}
}
