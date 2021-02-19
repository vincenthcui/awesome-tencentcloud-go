package tencentcloud

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"time"
)

const (
	codeRequestLimitExceeded = "RequestLimitExceeded"
)

// OnRateLimitExceeded will retry on rate limit exceeded in 1s
func OnRateLimitExceeded(ctx context.Context, action Action, request, response interface{}, err error) error {
	if err != nil {
		return err
	}
	if err, ok := err.(*errors.TencentCloudSDKError); ok {
		if err.Code == codeRequestLimitExceeded {
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
