package recursiveinternal

import "os"

func isDir(info os.FileInfo, err error) bool {
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
