package fs

import "os"

func IsNotPathExistsUsing(fileInfo os.FileInfo, err error) bool {
	return os.IsNotExist(err) || fileInfo == nil
}
