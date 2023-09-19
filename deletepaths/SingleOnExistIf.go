package deletepaths

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func SingleOnExistIf(
	isRemove bool,
	location string,
) *errorwrapper.Wrapper {
	if !isRemove {
		return nil
	}

	if !fsinternal.IsPathExists(location) {
		return nil
	}

	err := os.Remove(location)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(errtype.DeletePathFailed, err, location)
}
