package fs

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func ReadStringOnExistUsingLock(filePath string) *errstr.Result {
	if IsPathExistsUsingLock(filePath) {
		return ReadFileStringUsingLock(filePath)
	}

	return &errstr.Result{
		Value:        constants.EmptyString,
		ErrorWrapper: nil,
	}
}
