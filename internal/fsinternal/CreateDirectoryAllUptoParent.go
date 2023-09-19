package fsinternal

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
)

func CreateDirectoryAllUptoParent(
	location string,
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	parentDir := ParentDir(location)

	if IsDirectory(parentDir) {
		return nil
	}

	return CreateDirectoryAll(parentDir, fileMode)
}
