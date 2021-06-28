package createpath

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CreateSingleUsingFileMode(
	mode os.FileMode,
	filePath string,
) (
	*os.File,
	*errorwrapper.Wrapper,
) {
	file, err := os.Create(filePath)

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
