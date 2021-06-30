package fs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ReadFileStringIfExistUsingLock(filePath string) *errstr.Result {
	if IsPathExistsUsingLock(filePath) {
		return ReadFileStringUsingLock(filePath)
	}

	return &errstr.Result{
		Value: constants.EmptyString,
		ErrorWrapper: errnew.PathMessages(
			errtype.PathMissingOrInvalid,
			filePath),
	}
}
