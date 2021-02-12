package tencentcloud_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"
	"testing"
)

const (
	secretID  = ""
	secretKey = ""
)

func TestClient_Send(t *testing.T) {
	cli := tencentcloud.NewClient(tencentcloud.WithSecret(secretID, secretKey))
	req := cvm.DescribeInstancesRequest{}
	resp := cvm.DescribeInstancesResponse{}
	err := cli.Send(context.Background(), tencentcloud.CVMDescribeInstances, &req, &resp)
	assert.NoError(t, err)
}
