package createpath

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CreateSingle(
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

	return file, errnew.EmptyPtr
}
