package pathcompareinternal

import (
	"time"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/corecomparator"
)

func LastModified(
	left, right time.Time,
) corecomparator.Compare {
	return corecmp.Time(left, right)
}
