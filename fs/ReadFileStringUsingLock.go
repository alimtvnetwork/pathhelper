package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func ReadFileStringUsingLock(filePath string) *errstr.Result {
	errBytes := ReadFileUsingLock(filePath)

	return &errstr.Result{
		Value:        errBytes.String(),
		ErrorWrapper: errBytes.ErrorWrapper,
	}
}
