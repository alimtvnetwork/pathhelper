package pathhelper

import "gitlab.com/evatix-go/pathhelper/fileinfo"

func getFileInfos(path string) *fileinfo.Wrappers {
	return fileinfo.NewWrappersPtr(path)
}
