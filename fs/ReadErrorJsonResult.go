package fs

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
)

func ReadErrorJsonResult(filePath string) *errjson.Result {
	errBytes := ReadFile(filePath)

	if errBytes.IsFailed() {
		return errjson.EmptyWithErrorPtrUsingErrorWrapper(
			errBytes.ErrorWrapper)
	}

	return errjson.NewBytesPtr(
		errBytes.Values,
		errBytes.ErrorWrapper)
}
