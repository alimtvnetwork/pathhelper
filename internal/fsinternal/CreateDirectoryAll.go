package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CreateDirectoryAll(
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	return errnew.Path(
		errtype.CreateDirectoryFailed,
		os.MkdirAll(location, mode),
		location)
}
