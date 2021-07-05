package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteStringToFileUsingLockFileMode(
	filePath string,
	content string,
	fileMode os.FileMode,
	isKeepExistingFileModeOnExist bool,
) *errorwrapper.Wrapper {
	return WriteFileUsingFileMode(
		filePath,
		[]byte(content),
		fileMode,
		isKeepExistingFileModeOnExist)
}
