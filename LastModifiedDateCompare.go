package pathhelper

import (
	"time"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
)

func LastModifiedDateCompare(
	left, right *time.Time,
) corecomparator.Compare {
	return corecmp.TimePtr(left, right)
}
