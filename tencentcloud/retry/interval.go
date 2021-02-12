package retry

import (
	"math"
	"time"
)

// retry interval
// return 0 if times equals 0
type Interval func(times int) time.Duration

func ExponentialBackoff(times int) time.Duration {
	if times == 0 {
		return 0
	}

	seconds := int64(math.Pow(2, float64(times)))
	return time.Duration(seconds) * time.Second
}
