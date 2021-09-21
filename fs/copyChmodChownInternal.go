package fs

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func copyChmodChownInternal(
	srcPath,
	dstPath string,
	sourceFileInfo os.FileInfo,
) *errorwrapper.Wrapper {
	err := os.Chmod(dstPath, sourceFileInfo.Mode())

	if err != nil {
		return errnew.SourceDestinationMessages(
			errtype.ChmodApplyFailed,
			srcPath,
			dstPath,
			err.Error())
	}

	if osconsts.IsLinux {
		srcInstance := fileinfopath.Instance{
			FileInfo: sourceFileInfo,
			FullPath: srcPath,
			Error:    nil,
		}

		return pathsysinfo.ChownCopyUsing(
			&srcInstance,
			dstPath)
	}

	return errnew.EmptyPtr
}
