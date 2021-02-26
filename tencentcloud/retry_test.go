package tencentcloud_test

import (
	"context"
	act "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions/cvm/2017-03-12"
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
	err := client.Send(ctx, act.RunInstances, req, resp)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}
