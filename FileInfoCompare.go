package pathhelper

import (
	"os"

	"gitlab.com/auk-go/core/corecomparator"
	"gitlab.com/auk-go/pathhelper/internal/pathcompareinternal"
)

func FileInfoCompare(
	left, right os.FileInfo,
) corecomparator.Compare {
	return pathcompareinternal.FileInfoLastModified(left, right)
}
