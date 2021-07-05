package fs

import (
	"fmt"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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

	dstFileInfo, dstErr := os.Stat(dstPath)
	if IsNotPathExistsUsing(dstFileInfo, dstErr) {
		return errnew.Path(
			errtype.PathStatFailed,
			dstErr,
			dstPath)
	}

	if !(dstFileInfo.Mode().IsRegular()) {
		cannotCopyErr := fmt.Errorf(
			"CopyFile: non-regular destination file %s (%q)",
			dstFileInfo.Name(),
			dstFileInfo.Mode().String())

		return errnew.Path(
			errtype.Copy,
			cannotCopyErr,
			srcPath)
	}

	if os.SameFile(sourceFileInfo, dstFileInfo) {
		return errnew.EmptyPtr
	}

	linkErr := os.Link(srcPath, dstPath)

	if linkErr == nil {
		return errnew.EmptyPtr
	}

	return CopyFileContents(srcPath, dstPath)
}
