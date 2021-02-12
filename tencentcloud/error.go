package tencentcloud

import (
	"fmt"
	"time"
)

type RetryException struct {
	In time.Duration
}

func (retry RetryException) Error() string {
	return fmt.Sprintf("retry in %f second", retry.In.Seconds())
}
