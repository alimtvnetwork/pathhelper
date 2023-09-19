package fsinternal

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

// MoveFile move file path from source to destination
func MoveFile(srcPath, dstPath string) *errorwrapper.Wrapper {
	err := os.Rename(srcPath, dstPath)

	return errnew.
		Path.
		Error(
			errtype.FileOrDirectoryRelatedExecution,
			err,
			srcPath)
}
