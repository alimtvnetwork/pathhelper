package ds

import "os"

func isDir(location string) bool {
	fileInfo, err := os.Stat(location)

	if os.IsNotExist(err) {
		return false
	}

	return fileInfo.IsDir()
}
