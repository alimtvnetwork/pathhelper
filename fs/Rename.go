package fs

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func Rename(srcPath, dstPath string) *errorwrapper.Wrapper {
	err := os.Rename(srcPath, dstPath)

	return errnew.
		Path.
		Error(
			errtype.RenamePathFailed,
			err,
			srcPath)
}
