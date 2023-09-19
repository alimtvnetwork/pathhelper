package fsinternal

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func CreateDirectoryAll(
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	err := os.MkdirAll(location, mode)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			err,
			location)
}
