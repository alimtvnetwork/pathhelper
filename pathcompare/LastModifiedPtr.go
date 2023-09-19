package pathcompare

import (
	"time"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
)

func LastModifiedPtr(
	left, right *time.Time,
) corecomparator.Compare {
	return corecmp.TimePtr(left, right)
}
