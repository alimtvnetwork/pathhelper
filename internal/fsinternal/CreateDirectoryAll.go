package fsinternal

import (
	"os"

	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func CreateDirectoryAll(path string) *errorwrapper.Wrapper {
	// TODO: we need to think about default permission
	// https://gitlab.com/evatix-go/nginxconf/-/issues/53
	return errnew.ErrPtr(os.MkdirAll(path, filemode.X644))
}
