package fs

import "os"

func IsNotPathExistsUsing(fileInfo os.FileInfo, err error) bool {
	isExist := err == nil || !os.IsNotExist(err)

	return !isExist || fileInfo == nil
}
