package pathhelper

import "os"

func IsExistButNotDirectory(location string) bool {
	fileInfo, err := os.Stat(location)
	isExist := err == nil || !os.IsNotExist(err)

	return isExist &&
		fileInfo != nil &&
		fileInfo.IsDir()
}
