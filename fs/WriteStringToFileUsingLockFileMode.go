package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringToFileUsingLockFileMode(
	isCreateParentDir,
	isKeepExistingFileModeOnExist bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	return WriteFileUsingFileMode(
		isCreateParentDir,
		isKeepExistingFileModeOnExist,
		dirMode,
		fileMode,
		filePath,
		[]byte(content),
	)
}
