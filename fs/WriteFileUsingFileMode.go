package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteFileUsingFileMode(
	isCreateParentDir,
	isKeepExistingFileModeOnExist bool,
	filePath string,
	content []byte,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	if content == nil {
		return errnew.Messages.Many(
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
		return writeExistingFileContentUsingFileMode(
			filePath,
			content,
			mode)
	}

	// new content
	return writeNewFileContentUsingFileMode(
		isCreateParentDir,
		filePath,
		content,
		mode)
}
