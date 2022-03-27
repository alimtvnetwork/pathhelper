package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func JsonWriteMarshalUsingLock(
	isCreateParentDir,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	marshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return JsonWriteMarshal(
		isCreateParentDir,
		isSkipOnNilObject,
		isKeepExistingFileModeOnExist,
		dirMode,
		fileMode,
		filePath,
		marshallingObjectRef)
}
