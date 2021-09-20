package fs

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
)

// CopyFile Future ref: https://stackoverflow.com/a/21067803
func CopyFile(srcPath, dstPath string) *errorwrapper.Wrapper {
	sourceFileInfo, err := os.Stat(srcPath)
	if IsNotPathExistsUsing(sourceFileInfo, err) {
		return errnew.Path(
			errtype.PathStatFailed,
			err,
			srcPath)
	}

	if !sourceFileInfo.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		cannotCopySymLinkErr := fmt.Errorf(
			"CopyFile: non-regular source file %s (%q)",
			sourceFileInfo.Name(),
			sourceFileInfo.Mode().String())

		return errnew.Path(
			errtype.Copy,
			cannotCopySymLinkErr,
			srcPath)
	}

	if sourceFileInfo.IsDir() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		cannotCopyDir := fmt.Errorf(
			"CopyFile: don't support dir copy %s (%q)",
			sourceFileInfo.Name(),
			srcPath)

		return errnew.Path(
			errtype.Copy,
			cannotCopyDir,
			srcPath)
	}

	dstFileInfo, dstErr := os.Stat(dstPath)
	isExist := IsPathExistsUsing(dstFileInfo, dstErr)
	if isExist && !dstFileInfo.IsDir() {
		return deletepaths.Recursive(dstPath)
	} else if isExist && dstFileInfo.IsDir() {
		return errnew.PathMessages(
			errtype.PathCopy,
			dstPath,
			"don't support copy dir on file copier. destination contains same file name dir.")
	}

	// copy new file
	return CopyFileContents(srcPath, dstPath)
}
