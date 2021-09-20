package pathsysinfo

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func ChownCopyUnix(srcFullPath, dstFullPath string) *errorwrapper.Wrapper {
	if osconsts.IsWindows {
		return errnew.EmptyPtr
	}

	return ChownCopy(srcFullPath, dstFullPath)
}
