package createpath

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateSingleUsingFileMode(
	mode os.FileMode,
	filePath string,
) (
	*os.File,
	*errorwrapper.Wrapper,
) {
	dirCreateErr := fsinternal.CreateDirectoryAllUptoParent(
		filePath)

	if dirCreateErr.HasError() {
		return nil, dirCreateErr
	}

	file, err := os.Create(
		filePath)

	if err != nil {
		return file, errnew.Path(
			errtype.CreatePathFailed,
			err,
			filePath)
	}

	chmodErr := os.Chmod(filePath, mode)

	if chmodErr != nil {
		return file, errnew.Path(
			errtype.ChmodApplyFailed,
			err,
			filePath)
	}

	return file, errnew.EmptyPtr
}
