package tencentcloud

import (
	"context"
	"fmt"
	"net"
	"time"
)

// raise NeedRetry in interceptor will retry current request in In seconds
type NeedRetry struct {
	In time.Duration
}

func (retry NeedRetry) Error() string {
	return fmt.Sprintf("retry in %f second", retry.In.Seconds())
}

// retry on temporary or timeout network failure
func OnNetworkFailure(ctx context.Context, action Action, request, response interface{}, err error) error {
	if err, ok := err.(net.Error); ok {
		if err.Temporary() || err.Timeout() {
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
