package pathhelper

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompare"
)

func SizeCompare(
	left, right *int64,
) corecomparator.Compare {
	return pathcompare.Size(left, right)
}
