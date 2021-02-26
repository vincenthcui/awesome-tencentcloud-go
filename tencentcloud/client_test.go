package tencentcloud_test

import (
	"context"
	act "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions/cvm/2017-03-12"
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

	req := &cvm.DescribeInstancesRequest{}
	resp := &cvm.DescribeInstancesResponse{}

	cli := tencentcloud.NewClient(tencentcloud.WithSecret(secretID, secretKey))
	err := cli.Send(context.Background(), act.DescribeInstances, req, resp)
	assert.NoError(t, err)
}

func setup() {
	once.Do(func() {
		secretID, _ = os.LookupEnv("TC_SECRET_ID")
		secretKey, _ = os.LookupEnv("TC_SECRET_KEY")
	})
}
