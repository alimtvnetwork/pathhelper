package createpath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateManyUsingFileMode(
	mode os.FileMode,
	isIgnoreOnExist bool,
	files []string,
) ([]*os.File, *errorwrapper.Wrapper) {
	if len(files) == 0 {
		return []*os.File{}, errnew.EmptyPtr
	}

	slice := make(
		[]*os.File,
		constants.Zero,
		len(files))

	if isIgnoreOnExist {
		for _, filePath := range files {
			if fsinternal.IsPathExists(filePath) {
				continue
			}

			file, errWp := CreateSingleUsingFileMode(
				mode,
				filePath,
			)

			if errWp.HasError() {
				return slice, errWp
			}

			slice = append(slice, file)
		}

		return slice, errnew.EmptyPtr
	}

	// no checking create
	for _, filePath := range files {
		file, errWp := CreateSingleUsingFileMode(
			mode,
			filePath,
		)

		if errWp.HasError() {
			return slice, errWp
		}

		slice = append(slice, file)
	}

	return slice, errnew.EmptyPtr
}
