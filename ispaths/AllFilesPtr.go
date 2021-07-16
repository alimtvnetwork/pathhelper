package ispaths

import (
	"gitlab.com/evatix-go/pathhelper/internal/fileinfogetter"
)

func AllFilesPtr(fullPaths *[]string) bool {
	if fullPaths == nil {
		return false
	}

	convertedFileInfos := fileinfogetter.Get(
		fullPaths)

	for _, wrapper := range *convertedFileInfos {
		isAnyInvalidOrDir := wrapper == nil ||
			wrapper.IsDir()

		if isAnyInvalidOrDir {
			return false
		}
	}

	return true
}
