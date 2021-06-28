package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func CreateDirectoryAllUptoParent(location string) *errorwrapper.Wrapper {
	parentDir := ParentDir(location)

	if IsDirectory(parentDir) {
		return errnew.EmptyPtr
	}

	return CreateDirectoryAllDefault(parentDir)
}
