package deletepaths

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func SingleOnExist(location string) *errorwrapper.Wrapper {
	if !fsinternal.IsPathExists(location) {
		return errnew.EmptyPtr
	}

	err := os.Remove(location)

	return errnew.Path(errtype.DeletePathFailed, err, location)
}
