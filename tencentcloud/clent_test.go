package tencentcloud_test

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"
)

var (
	secretID  = ""
	secretKey = ""
	once      sync.Once
)

func TestClient_Send(t *testing.T) {
	setup()

	cli := tencentcloud.NewClient(tencentcloud.WithSecret(secretID, secretKey))
	req := &cvm.DescribeInstancesRequest{}
	resp := &cvm.DescribeInstancesResponse{}
	err := cli.Send(context.Background(), tencentcloud.CVMDescribeInstances, req, resp)
	assert.NoError(t, err)
}

func setup() {
	once.Do(func() {
		var ok bool
		if secretID, ok = os.LookupEnv("TC_SECRET_ID"); !ok {
			panic("env TC_SECRET_ID is required for test")
		}
		if secretKey, _ = os.LookupEnv("TC_SECRET_KEY"); !ok {
			panic("env TC_SECRET_KEY is required for test")
		}
	})
}
