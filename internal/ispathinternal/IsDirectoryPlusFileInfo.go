package ispathinternal

import "os"

func IsDirectoryPlusFileInfo(location string) (isSuccess bool, fileInfo os.FileInfo) {
	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) {
		return false, fileInfo
	}

	return fileInfo.IsDir(), fileInfo
}
