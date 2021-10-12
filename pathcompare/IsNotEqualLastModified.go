package pathcompare

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func IsNotEqualLastModified(
	left, right time.Time,
) bool {
	return !corecomparator.Time(left, right).IsEqual()
}
