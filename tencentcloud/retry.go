package tencentcloud

import (
	"context"
	"net"
	"time"
)

// retry on temporary or timeout network failure
func OnNetworkFailure(ctx context.Context, action Action, request, response interface{}, err error) error {
	if err, ok := err.(net.Error); ok {
		if err.Temporary() || err.Timeout() {
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
