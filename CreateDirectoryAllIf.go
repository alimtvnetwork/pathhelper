package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/dirinfo"
)

func CreateDirectoryAllIf(
	condition bool,
	path string,
	fileMode os.FileMode,
) *fileinfo.Result {
	if condition {
		return CreateDirectoryAll(path, fileMode)
	}

	return fileinfo.NewEmptyDirectoryResult()
}
