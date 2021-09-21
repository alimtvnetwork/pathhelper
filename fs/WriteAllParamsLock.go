package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func WriteAllParamsLock(
	isCreateParentDir,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	fileMod os.FileMode,
	dirCreateMod os.FileMode,
	filePath string,
	contents []byte,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteAllParams(
		isCreateParentDir,
		isSkipOnNilObject,
		isKeepExistingFileModeOnExist,
		fileMod,
		dirCreateMod,
		filePath,
		contents)
}
