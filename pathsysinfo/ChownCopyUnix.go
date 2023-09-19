package pathsysinfo

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
)

func ChownCopyUnix(srcFullPath, dstFullPath string) *errorwrapper.Wrapper {
	if osconsts.IsWindows {
		return nil
	}

	return ChownCopy(srcFullPath, dstFullPath)
}
