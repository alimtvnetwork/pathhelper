package ispath

import "gitlab.com/evatix-go/pathhelper"

func Exist(path string) bool {
	fileInfoWrapper := pathhelper.GetFileInfoWrapper(path)

	return fileInfoWrapper.IsPathExists()
}
