package fs

import (
	"os"

	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/fileinfopath"
	"gitlab.com/auk-go/pathhelper/pathsysinfo"
)

func copyChmodChownInternal(
	srcPath,
	dstPath string,
	sourceFileInfo os.FileInfo,
) *errorwrapper.Wrapper {
	err := os.Chmod(dstPath, sourceFileInfo.Mode())

	if err != nil {
		return errnew.SrcDst.Error(
			errtype.ChmodApplyFailed,
			err,
			srcPath,
			dstPath,
		)
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

	return nil
}
