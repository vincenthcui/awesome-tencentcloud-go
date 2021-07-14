package tencentcloud

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/metrics"
)

// raise NeedRetry in interceptor will retry current request in In seconds
type NeedRetry struct {
	In time.Duration
}

func (retry NeedRetry) Error() string {
	return fmt.Sprintf("retry in %f second", retry.In.Seconds())
}

// retry on temporary or timeout network failure
func OnNetworkFailure(ctx context.Context, act Action, request, response interface{}, err error) error {
	if err, ok := err.(net.Error); ok {
		if err.Temporary() || err.Timeout() {
			metrics.NetworkFailureRetryTotal.WithLabelValues(act.Service, act.Action, act.Version).Add(1)
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
