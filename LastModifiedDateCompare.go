package pathhelper

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func LastModifiedDateCompare(
	left, right *time.Time,
) corecomparator.Compare {
	return corecomparator.TimePtr(left, right)
}
