package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteAllParamsLock(
	isCreateParentDir,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirCreateMode os.FileMode,
	fileMode os.FileMode,
	filePath string,
	contents []byte,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteAllParams(
		isCreateParentDir,
		isSkipOnNilObject,
		isKeepExistingFileModeOnExist,
		dirCreateMode,
		fileMode,
		filePath,
		contents)
}
