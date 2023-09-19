package createdir

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func AllUptoParentDefault(location string) *errorwrapper.Wrapper {
	return fsinternal.CreateDirectoryAllUptoParent(location, DefaultDirectoryFileMode)
}
