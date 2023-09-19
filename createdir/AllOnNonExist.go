package createdir

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func AllOnNonExist(
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	if fsinternal.IsExistButDirectory(location) {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			os.MkdirAll(location, mode),
			location)
}
