package tencentcloud

import (
	"fmt"
	"time"
)

type NeedRetry struct {
	In time.Duration
}

func (retry NeedRetry) Error() string {
	return fmt.Sprintf("retry in %f second", retry.In.Seconds())
}
