package pathcompareinternal

import (
	"time"

	"gitlab.com/auk-go/core/corecmp"
)

func IsLastModifiedEqual(
	left, right time.Time,
) bool {
	return corecmp.Time(left, right).IsEqual()
}
