package tencentcloud

import (
	"encoding/json"

	terrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	common "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

func maybeFailed(bytes []byte) *terrors.TencentCloudSDKError {
	payload := common.ErrorResponse{}
	err := json.Unmarshal(bytes, &payload)
	if err != nil {
		return &terrors.TencentCloudSDKError{
			Code:    "ClientError.ParseJsonError",
			Message: err.Error(),
		}
	}
	if payload.Response.Error.Code != "" {
		return &terrors.TencentCloudSDKError{
			Code:      payload.Response.Error.Code,
			Message:   payload.Response.Error.Message,
			RequestId: payload.Response.RequestId,
		}
	}
	return nil
}
