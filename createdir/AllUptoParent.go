package createdir

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func AllUptoParent(location string, mode os.FileMode) *errorwrapper.Wrapper {
	return fsinternal.CreateDirectoryAllUptoParent(location, mode)
}
