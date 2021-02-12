package tencentcloud

import (
	"context"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions"
	"net"
	"time"
)

// retry on temporary or timeout network failure
func OnNetworkFailure(ctx context.Context, action actions.Action, request, response interface{}, err error) error {
	if err, ok := err.(net.Error); ok {
		if err.Temporary() || err.Timeout() {
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
