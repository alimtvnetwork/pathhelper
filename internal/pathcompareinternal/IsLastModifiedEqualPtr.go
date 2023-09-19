package pathcompareinternal

import (
	"time"

	"gitlab.com/auk-go/core/corecmp"
)

func IsLastModifiedEqualPtr(
	left, right *time.Time,
) bool {
	return corecmp.TimePtr(left, right).IsEqual()
}
