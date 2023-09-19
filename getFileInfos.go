package pathhelper

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/fileinfo"
)

func GetFileInfoWrappersFrom(path string, isNormalize bool) *fileinfo.Wrappers {
	return fileinfo.NewWrappersPtr(
		path,
		osconsts.PathSeparator,
		isNormalize)
}
