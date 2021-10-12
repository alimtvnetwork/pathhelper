package pathcompareinternal

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func IsLastModifiedEqualPtr(
	left, right *time.Time,
) bool {
	return corecomparator.TimePtr(left, right).IsEqual()
}
