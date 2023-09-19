package deletepaths

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func Recursive(location string) *errorwrapper.Wrapper {
	err := os.RemoveAll(location)
	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.DeletePathFailed,
			err,
			location+"->recursive remove failed.")
}
