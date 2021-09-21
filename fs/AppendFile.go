package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func AppendFile(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			"Cannot write empty or nil contents to the file.",
			filePath)
	}

	// file already exist, append
	if IsPathExists(filePath) {
		return appendFileContent(filePath, content)
	}

	return writeNewFileContent(
		isCreateParentDir,
		filePath,
		content)
}
