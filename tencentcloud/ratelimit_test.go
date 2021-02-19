package tencentcloud_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	tc "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"
)

type limitAlwaysExceeded struct{}

func (c *limitAlwaysExceeded) Do(*http.Request) (*http.Response, error) {
	resp := []byte(`{"Response":{"Error": {"Code":"RequestLimitExceeded"}}}`)
	return &http.Response{
		StatusCode:    http.StatusOK,
		Body:          ioutil.NopCloser(bytes.NewBuffer(resp)),
		ContentLength: int64(len(resp)),
	}, nil
}

func TestOnRateLimitExceeded(t *testing.T) {
	setup()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	req := cvm.NewRunInstancesRequest()
	resp := cvm.NewRunInstancesResponse()

	client := tc.NewClient(tc.WithSecret(secretID, secretID), tc.WithHttpClient(&limitAlwaysExceeded{}))
	err := client.Send(ctx, tc.CVMRunInstances, req, resp)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}
