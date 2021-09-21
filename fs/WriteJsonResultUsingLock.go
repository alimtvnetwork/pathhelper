package fs

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteJsonResultUsingLock(
	isCreateParentDir,
	isSkipErrorOnNilOrEmpty bool,
	jsonResult *corejson.Result,
	location string,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteJsonResult(
		isCreateParentDir,
		isSkipErrorOnNilOrEmpty,
		jsonResult,
		location)
}
