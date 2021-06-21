package createdirinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

// AllRecurse Create all sub-directories and create the final directory
func AllRecurse(
	path string,
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	isIgnoredAction := fsinternal.IsPathExists(path)

	if !isIgnoredAction {
		err := os.MkdirAll(path, fileMode)

		return errorwrapper.NewFilePtr(errtype.Directory, err, path)
	}

	return errnew.EmptyPtr
}
