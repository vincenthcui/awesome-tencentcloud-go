package tencentcloud

import (
	"context"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"

	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/metrics"
)

const (
	codeRequestLimitExceeded = "RequestLimitExceeded"
)

// OnRateLimitExceeded will retry on rate limit exceeded in 1s
func OnRateLimitExceeded(ctx context.Context, act Action, request, response interface{}, err error) error {
	if err, ok := err.(*errors.TencentCloudSDKError); ok {
		if err.Code == codeRequestLimitExceeded {
			metrics.RateLimitRetryTotal.WithLabelValues(act.Service, act.Action, act.Version).Add(1)
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
