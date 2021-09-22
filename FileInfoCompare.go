package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompare"
)

func FileInfoCompare(
	left, right os.FileInfo,
) corecomparator.Compare {
	return pathcompare.FileInfo(left, right)
}
