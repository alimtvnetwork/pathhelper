package deletepaths

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func RecursiveOnExist(location string) *errorwrapper.Wrapper {
	if !fsinternal.IsPathExists(location) || len(location) == 0 {
		return errnew.EmptyPtr
	}

	err := os.RemoveAll(location)

	return errnew.Path(
		errtype.DeletePathFailed,
		err,
		location+"->recursive remove failed.")
}
