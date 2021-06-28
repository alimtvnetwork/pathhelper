package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// CopyFile Future ref: https://stackoverflow.com/a/21067803
func CopyFile(srcPath, dstPath string) *errorwrapper.Wrapper {
	err := copyFileContents(srcPath, dstPath)

	return errnew.Path(
		errtype.FileOrDirectoryRelatedExecution,
		err,
		srcPath)
}
