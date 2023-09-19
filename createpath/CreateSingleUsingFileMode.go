package createpath

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
)

func CreateSingleUsingFileMode(
	isLock bool,
	mode os.FileMode,
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

	file, err := os.Create(
		filePath)

	if err != nil {
		return file, errnew.
			Path.
			Error(
				errtype.CreatePathFailed,
				err,
				filePath)
	}

	chmodErr := os.Chmod(filePath, mode)

	if chmodErr != nil {
		return file, errnew.
			Path.
			Error(
				errtype.ChmodApplyFailed,
				err,
				filePath)
	}

	return file, nil
}
