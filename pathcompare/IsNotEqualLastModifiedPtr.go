package pathcompare

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func IsNotEqualLastModifiedPtr(
	left, right *time.Time,
) bool {
	return !corecomparator.TimePtr(left, right).IsEqual()
}
