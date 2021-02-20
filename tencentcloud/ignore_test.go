package tencentcloud

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreResponse(t *testing.T) {
	bytes := []byte(`{"hello": "world"}`)
	err := json.Unmarshal(bytes, IgnoreResponse())
	assert.NoError(t, err)
}
