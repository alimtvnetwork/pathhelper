package createdirinternal

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

// AllRecurse Create all sub-directories and create the final directory
func AllRecurse(
	path string,
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	isIgnoredAction := fsinternal.IsPathExists(path)

	if isIgnoredAction {
		return nil
	}

	err := os.MkdirAll(path, fileMode)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.Directory,
			err,
			path)
}
