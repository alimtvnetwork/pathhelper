package pathcompareinternal

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func LastModified(
	left, right time.Time,
) corecomparator.Compare {
	return corecomparator.Time(left, right)
}
