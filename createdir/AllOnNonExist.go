package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AllOnNonExist(
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	if fsinternal.IsExistButDirectory(location) {
		return errnew.EmptyPtr
	}

	return errnew.Path(
		errtype.CreateDirectoryFailed,
		os.MkdirAll(location, mode),
		location)
}
