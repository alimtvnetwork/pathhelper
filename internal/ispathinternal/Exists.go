package ispathinternal

import "gitlab.com/evatix-go/pathhelper/fileinfo"

func Exists(path string) bool {
	fileInfoWrapper := fileinfo.New(path)

	return fileInfoWrapper.IsPathExists()
}
