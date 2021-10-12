package pathcompareinternal

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
)

func IsLastModifiedEqual(
	left, right time.Time,
) bool {
	return corecomparator.Time(left, right).IsEqual()
}
