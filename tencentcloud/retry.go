package tencentcloud

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions"
	"net"
	"time"
)

const (
	CodeRequestLimitExceeded = "CodeRequestLimitExceeded"
)

// retry on temporary or timeout network failure
func OnNetworkFailure(ctx context.Context, action actions.Action, request, response interface{}, err error) error {
	if err != nil {
		return err
	}
	if err, ok := err.(net.Error); ok {
		if err.Temporary() || err.Timeout() {
			return NeedRetry{In: time.Second}
		}
	}
	return err
}

func OnRequestLimitExceeded(ctx context.Context, action actions.Action, request, response interface{}, err error) error {
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
