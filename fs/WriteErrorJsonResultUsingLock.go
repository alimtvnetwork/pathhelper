package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
)

func WriteErrorJsonResultUsingLock(
	isSkipErrorOnNilOrEmpty bool,
	errJsonResult *errjson.Result,
	location string,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteErrorJsonResult(
		isSkipErrorOnNilOrEmpty,
		errJsonResult,
		location,
	)
}
