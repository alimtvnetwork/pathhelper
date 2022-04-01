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
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return JsonWriteMarshal(
		isCreateParentDir,
		isSkipOnNilObject,
		isKeepExistingFileModeOnExist,
		dirMode,
		fileMode,
		filePath,
		marshallingObjectRef)
}
