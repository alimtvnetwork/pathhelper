package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func CreateDirectoryAllIf(
	condition bool,
	path string,
	fileMode os.FileMode,
) *pathhelpercore.DirectoryResult {
	if condition {
		return CreateDirectoryAll(path, fileMode)
	}

	return pathhelpercore.NewEmptyDirectoryResult()
}
