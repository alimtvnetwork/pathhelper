package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AllUptoParentMany(mode os.FileMode, locations ...string) *errorwrapper.Wrapper {
	for _, path := range locations {
		if errW := fsinternal.CreateDirectoryAllUptoParent(path, mode); errW.HasError() {
			return errW
		}
	}

	return errnew.EmptyPtr
}
