package fsinternal

import (
	"os"
)

func IsDirectory(location string) bool {
	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) {
		return false
	}

	return fileInfo != nil && fileInfo.IsDir()
}

func IsExistButNotDirectory(location string) bool {
	fileInfo, err := os.Stat(location)
	isExist := err == nil || !os.IsNotExist(err)

	return isExist &&
		fileInfo != nil &&
		fileInfo.IsDir()
}
