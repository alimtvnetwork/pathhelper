package pathhelper

import (
	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/pathhelper/internal/pathcompareinternal"
)

func SizeCompare(
	left, right *int64,
) corecomparator.Compare {
	return pathcompareinternal.SizePtr(left, right)
}
