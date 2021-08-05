package fsinternal

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func CreateDirectoryAllUptoParentDefault(location string) *errorwrapper.Wrapper {
	parentDir := ParentDir(location)

	if IsDirectory(parentDir) {
		return errnew.EmptyPtr
	}

	return CreateDirectoryAllDefault(parentDir)
}
