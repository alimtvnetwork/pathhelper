package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func RemoveCreateAll(
	isLock,
	isSkipOnExist,
	isRemove,
	isApplyCreate bool,
	location string,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	if isLock {
		mutexLock.Lock()
		defer mutexLock.Unlock()
	}

	if isSkipOnExist && fsinternal.IsPathExists(location) {
		return errnew.EmptyPtr
	}

	var removeErr error
	if isRemove {
		removeErr = os.RemoveAll(location)
	}

	if removeErr != nil {
		return errnew.Path(
			errtype.RemoveFailed,
			removeErr,
			location)
	}

	if isApplyCreate {
		return errnew.Path(
			errtype.CreateDirectoryFailed,
			os.MkdirAll(location, mode),
			location)
	}

	return errnew.EmptyPtr
}
