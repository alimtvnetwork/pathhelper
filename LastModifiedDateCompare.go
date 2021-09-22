package pathhelper

import (
	"time"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompare"
)

func LastModifiedDateCompare(
	left, right *time.Time,
) corecomparator.Compare {
	return pathcompare.LastModified(left, right)
}
