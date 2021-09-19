package fs

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/createdir"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

func CreateParentDirWithChmodChown(srcPath string, dstPath string) *errorwrapper.Wrapper {
	srcDir := splitinternal.GetBaseDir(srcPath)
	srcBaseDirInfo, err := os.Stat(srcDir)

	if IsNotPathExistsUsing(srcBaseDirInfo, err) {
		return errnew.Path(errtype.Copy, err, srcDir)
	}

	createDirErr := createdir.AllUptoParent(
		dstPath,
		srcBaseDirInfo.Mode())

	if createDirErr.HasError() {
		return createDirErr
	}

	if osconsts.IsLinux {
		return pathsysinfo.ChownCopy(srcPath, dstPath)
	}

	return errnew.EmptyPtr
}
