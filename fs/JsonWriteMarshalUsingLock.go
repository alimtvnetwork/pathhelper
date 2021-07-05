package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func JsonWriteMarshalUsingLock(
	isSkipOnNilObject bool,
	filePath string,
	fileMod os.FileMode,
	isKeepExistingFileModeOnExist bool,
	marshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return JsonWriteMarshal(
		isSkipOnNilObject,
		filePath,
		fileMod,
		isKeepExistingFileModeOnExist,
		marshallingObjectRef)
}
