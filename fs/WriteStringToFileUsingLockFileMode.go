package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringToFileUsingLockFileMode(
	isCreateParentDir,
	isKeepExistingFileModeOnExist bool,
	filePath string,
	content string,
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	return WriteFileUsingFileMode(
		isCreateParentDir,
		isKeepExistingFileModeOnExist,
		filePath,
		[]byte(content),
		fileMode,
	)
}
