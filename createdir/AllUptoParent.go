package createdir

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func AllUptoParent(location string) *errorwrapper.Wrapper {
	return fsinternal.CreateDirectoryAllUptoParent(location)
}
