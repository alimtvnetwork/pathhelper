package fs

import "os"

func IsPathExistsUsing(fileInfo os.FileInfo, err error) bool {
	return !os.IsNotExist(err) || fileInfo != nil
}
