package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

func GetFileInfoWrappersFrom(path string) *fileinfo.Wrappers {
	return fileinfo.NewWrappersPtr(path)
}
