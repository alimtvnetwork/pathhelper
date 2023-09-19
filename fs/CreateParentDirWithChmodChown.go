package fs

import (
	"os"
	"path/filepath"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/createdir"
)

func CreateParentDirWithChmodChown(srcPath string, dstPath string) *errorwrapper.Wrapper {
	srcDir := filepath.Dir(srcPath)
	srcBaseDirInfo, err := os.Stat(srcDir)

	if IsNotPathExistsUsing(srcBaseDirInfo, err) {
		return errnew.
			Path.
			Error(
				errtype.Copy,
				err,
				srcDir)
	}

	dstParentDir := filepath.Dir(dstPath)
	createDirErr := createdir.AllOnNonExist(
		dstParentDir,
		srcBaseDirInfo.Mode())

	if createDirErr.HasError() {
		return createDirErr
	}

	return CopyChmodChown(srcDir, dstParentDir)
}
