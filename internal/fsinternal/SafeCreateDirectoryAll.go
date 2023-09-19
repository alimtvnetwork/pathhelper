package fsinternal

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func SafeCreateDirectoryAll(
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	if IsExistButDirectory(location) {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			os.MkdirAll(location, mode),
			location)
}
