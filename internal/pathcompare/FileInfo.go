package pathcompare

import (
	"os"

	"gitlab.com/evatix-go/core/corecomparator"
)

func FileInfo(
	left, right os.FileInfo,
) corecomparator.Compare {
	if left == nil && right == nil {
		return corecomparator.Equal
	}

	if left == nil || right == nil {
		return corecomparator.NotEqual
	}

	leftMod := left.ModTime()
	rightMod := right.ModTime()

	return LastModified(&leftMod, &rightMod)
}
