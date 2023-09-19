package fs

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errdefer"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/createdir"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func DirFileCreate(
	existingFileErrorWrapper *errorwrapper.Wrapper, // can be nil, current function return error
	directory,
	relativeFilePath string,
	chmod os.FileMode,
) *OsFile {
	dirCreateError := createdir.AllOnNonExist(
		directory,
		chmod)
	if dirCreateError.HasError() {
		return &OsFile{
			Location:         directory,
			OsFile:           nil,
			ErrorWrapper:     dirCreateError,
			DeferClosingFunc: nil,
		}
	}

	fullPath := pathjoin.JoinNormalized(
		directory,
		relativeFilePath)
	file, err := os.Create(fullPath)

	deferClosingFunc := func() *errorwrapper.Wrapper {
		return errdefer.CloseFile(
			fullPath,
			existingFileErrorWrapper,
			file)
	}

	if err != nil {
		return &OsFile{
			Location: fullPath,
			OsFile:   file,
			ErrorWrapper: errnew.
				Path.
				Error(
					errtype.CreatePathFailed,
					err,
					relativeFilePath),
			DeferClosingFunc: deferClosingFunc,
		}
	}

	chmodErr := file.Chmod(chmod)
	if chmodErr != nil {
		return &OsFile{
			Location: fullPath,
			OsFile:   file,
			ErrorWrapper: errnew.
				Path.
				Error(
					errtype.ChmodApplyFailed,
					chmodErr,
					relativeFilePath),
			DeferClosingFunc: deferClosingFunc,
		}
	}

	return &OsFile{
		Location:         fullPath,
		OsFile:           file,
		ErrorWrapper:     nil,
		DeferClosingFunc: deferClosingFunc,
	}
}
