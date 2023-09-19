package pathsysinfo

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/fileinfopath"
)

func ChownCopy(srcFullPath, dstFullPath string) *errorwrapper.Wrapper {
	srcInstance := fileinfopath.New(srcFullPath)

	return ChownCopyUsing(srcInstance, dstFullPath)
}
