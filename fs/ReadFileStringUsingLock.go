package fs

import "gitlab.com/auk-go/errorwrapper/errdata/errstr"

func ReadFileStringUsingLock(filePath string) *errstr.Result {
	errBytes := ReadFileUsingLock(filePath)

	return errBytes.ErrStr()
}
