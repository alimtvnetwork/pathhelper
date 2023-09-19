package pathcompare

import (
	"time"

	"gitlab.com/auk-go/core/corecmp"
)

func IsNotEqualLastModified(
	left, right time.Time,
) bool {
	return !corecmp.Time(left, right).IsEqual()
}
