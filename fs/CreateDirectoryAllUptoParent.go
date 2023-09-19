package fs

import (
	"gitlab.com/auk-go/errorwrapper"
)

func CreateDirectoryAllUptoParent(location string) *errorwrapper.Wrapper {
	parentDir := ParentDir(location)

	if IsDirectory(parentDir) {
		return nil
	}

	return CreateDirectoryAllDefault(parentDir)
}
