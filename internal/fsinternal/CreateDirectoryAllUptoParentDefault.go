package fsinternal

import (
	"gitlab.com/auk-go/errorwrapper"
)

func CreateDirectoryAllUptoParentDefault(location string) *errorwrapper.Wrapper {
	parentDir := ParentDir(location)

	if IsDirectory(parentDir) {
		return nil
	}

	return CreateDirectoryAllDefault(parentDir)
}
