package deletepaths

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func RecursiveOnExist(location string) *errorwrapper.Wrapper {
	if !fsinternal.IsPathExists(location) || len(location) == 0 {
		return nil
	}

	err := os.RemoveAll(location)
	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(errtype.DeletePathFailed,
			err,
			location+"->recursive remove failed.")
}
