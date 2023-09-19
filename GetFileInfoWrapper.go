package pathhelper

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/fileinfo"
)

func GetFileInfoWrapper(path string) *fileinfo.Wrapper {
	return fileinfo.New(path, osconsts.PathSeparator)
}
