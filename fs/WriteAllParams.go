package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/createdir"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func WriteAllParams(
	isCreateParentDir,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirCreateMod os.FileMode,
	fileMod os.FileMode,
	filePath string,
	contents []byte,
) *errorwrapper.Wrapper {
	if isSkipOnNilObject && contents == nil {
		return nil
	}

	if contents == nil {
		return fsinternal.NullContentErrorWrap(
			filePath)
	}

	isExist := IsPathExists(filePath)
	// file already exist
	if isExist && isKeepExistingFileModeOnExist {
		return writeExistingFileContent(
			filePath,
			contents)
	} else if isExist && !isKeepExistingFileModeOnExist {
		return writeExistingFileContentUsingFileMode(
			filePath,
			contents,
			fileMod)
	}

	// new
	var createDirErr *errorwrapper.Wrapper
	if isCreateParentDir {
		createDirErr = createdir.AllUptoParent(
			filePath,
			dirCreateMod)
	}

	if createDirErr.HasError() {
		return createDirErr
	}

	return WriteFile(
		false,
		filePath,
		contents)
}
