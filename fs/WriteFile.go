package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteFile(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			"Cannot write nil contents (bytes) to the file.",
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
