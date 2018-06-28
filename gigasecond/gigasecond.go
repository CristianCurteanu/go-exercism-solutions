package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return time.Unix(t.Unix()+int64(math.Pow(10, 9)), 0)
}
