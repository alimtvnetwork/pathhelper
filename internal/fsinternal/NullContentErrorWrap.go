package fsinternal

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func NullContentErrorWrap(filePath string) *errorwrapper.Wrapper {
	return errnew.Ref.Message(
		errtype.NullOrEmpty,
		"file-path",
		filePath,
		"cannot write or append empty or bytes into file.",
	)
}
