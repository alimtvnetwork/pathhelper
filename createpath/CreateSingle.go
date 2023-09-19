package createpath

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func CreateSingle(
	isLock bool,
	filePath string,
) (
	*os.File,
	*errorwrapper.Wrapper,
) {
	if isLock {
		lockerMutex.Lock()
		defer lockerMutex.Unlock()
	}

	dirCreateErr := fsinternal.CreateDirectoryAllUptoParentDefault(
		filePath)

	if dirCreateErr.HasError() {
		return nil, dirCreateErr
	}

	file, err := os.Create(filePath)

	if err != nil {
		return file, errnew.
			Path.
			Error(
				errtype.CreatePathFailed,
				err,
				filePath)
	}

	return file, nil
}
