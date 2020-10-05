package pathhelper

import "gitlab.com/evatix-go/pathhelper/pathhelpercore"

func GetFileInfoWrapper(path string) *pathhelpercore.FileInfoWrapper {
	return pathhelpercore.NewFileWrapperInfo(path)
}
