package pathcompare

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func LastModifiedPtr(
	left, right *time.Time,
) corecomparator.Compare {
	return corecomparator.TimePtr(left, right)
}
