package createpath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func CreateManySameDirWithFileMode(
	mode os.FileMode,
	isIgnoreOnExist bool,
	rootDir string,
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

			file, errWp := CreateSingle(
				filePath,
			)

			if errWp.HasError() {
				return slice, errWp
			}

			slice = append(slice, file)
		}

		return slice, applyLinuxRecursiveChmodOnDir(
			mode,
			rootDir)
	}

	// no checking create
	for _, filePath := range files {
		file, errWp := CreateSingle(
			filePath,
		)

		if errWp.HasError() {
			return slice, errWp
		}

		slice = append(slice, file)
	}

	return slice, applyLinuxRecursiveChmodOnDir(
		mode,
		rootDir)
}
