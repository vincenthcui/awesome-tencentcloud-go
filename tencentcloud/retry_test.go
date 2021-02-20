package tencentcloud_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	tc "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"
)

func TestOnNetworkFailure(t *testing.T) {
	setup()

	// nolint:govet
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	req := cvm.NewRunInstancesRequest()
	resp := cvm.NewRunInstancesResponse()

	client := tc.NewClient(tc.WithSecret(secretID, secretID), tc.WithHttpClient(&http.Client{
		Timeout: time.Duration(1),
	}))
	err := client.Send(ctx, tc.CVMRunInstances, req, resp)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}
