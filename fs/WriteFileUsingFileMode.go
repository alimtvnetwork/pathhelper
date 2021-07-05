package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteFileUsingFileMode(
	filePath string,
	content []byte,
	mode os.FileMode,
	isKeepExistingFileModeOnExist bool,
) *errorwrapper.Wrapper {
	if content == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			"Cannot write empty or nil contents to the file.",
			filePath)
	}

	isExist := IsPathExists(filePath)

	// file already exist
	if isExist && isKeepExistingFileModeOnExist {
		return writeExistingFileContent(
			filePath,
			content)
	} else if isExist && !isKeepExistingFileModeOnExist {
		return writeNewFileContentUsingFileMode(
			filePath,
			content,
			mode)
	}

	return writeNewFileContentUsingFileMode(
		filePath,
		content,
		mode)
}
