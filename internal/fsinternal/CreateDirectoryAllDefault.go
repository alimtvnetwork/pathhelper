package fsinternal

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/consts"
)

func CreateDirectoryAllDefault(location string) *errorwrapper.Wrapper {
	return errnew.
		Path.
		Error(
			errtype.CreateDirectoryFailed,
			os.MkdirAll(location, consts.DefaultDirectoryFileMode),
			location)
}
