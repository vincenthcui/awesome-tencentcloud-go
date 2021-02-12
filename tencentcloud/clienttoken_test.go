package tencentcloud

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"testing"
)

func TestInjectClientToken(t *testing.T) {
	request := &cvm.RunInstancesRequest{}
	injectClientToken(request)
	assert.NotNil(t, request.ClientToken)

	val, err := uuid.FromString(*request.ClientToken)
	assert.NoError(t, err)

	injectClientToken(request)
	assert.Equal(t, val.String(), *request.ClientToken)
}
