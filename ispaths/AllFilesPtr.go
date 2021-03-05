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
		if wrapper == nil || (*wrapper).IsDir() {
			return false
		}
	}

	return true
}
