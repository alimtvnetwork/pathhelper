package fs

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func ReadFileLinesUsingLock(filePath string) *errstr.Results {
	errString := ReadFileStringUsingLock(filePath)

	if errString.Value == "" {
		return errstr.New.Results.ErrorWrapper(
			errString.ErrorWrapper)
	}

	lines := strings.Split(
		errString.Value,
		constants.NewLineUnix)

	return &errstr.Results{
		Values:       lines,
		ErrorWrapper: errString.ErrorWrapper,
	}
}
