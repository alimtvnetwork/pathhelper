package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func WriteFile(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return fsinternal.NullContentErrorWrap(
			filePath)
	}

	// file already exist
	if IsPathExists(filePath) {
		return writeExistingFileContent(
			filePath,
			content)
	}

	return writeNewFileContent(
		isCreateParentDir,
		filePath,
		content)
}
