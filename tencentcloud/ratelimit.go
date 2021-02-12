package tencentcloud

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"time"
)

const (
	CodeRequestLimitExceeded = "CodeRequestLimitExceeded"
)

func OnRequestLimitExceeded(ctx context.Context, action Action, request, response interface{}, err error) error {
	if err != nil {
		return err
	}
	if err, ok := err.(*errors.TencentCloudSDKError); ok {
		if err.Code == CodeRequestLimitExceeded {
			return NeedRetry{In: time.Second}
		}
	}
	return err
}
