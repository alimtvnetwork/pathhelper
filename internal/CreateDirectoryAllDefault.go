package createdir

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateDirectoryAllDefault(location string) *errorwrapper.Wrapper {
	return fsinternal.CreateDirectoryAllDefault(
		location,
		consts.DefaultFileMode,
	)
}
