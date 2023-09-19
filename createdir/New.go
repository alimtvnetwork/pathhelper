package createdir

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/pathhelper"
	"gitlab.com/auk-go/pathhelper/dirinfo"
	"gitlab.com/auk-go/pathhelper/performingas"
)

// New Create directory and create the final directory
func New(path string, fileMode os.FileMode) *dirinfo.Result {
	fileInfoWrapper := pathhelper.GetFileInfoWrapper(path)
	isIgnoredAction := fileInfoWrapper.IsPathExists() || fileInfoWrapper.IsEmptyPath
	var errorWrapper *errorwrapper.Wrapper

	if !isIgnoredAction {
		err := os.MkdirAll(path, fileMode)
		errorWrapper = errnew.
			Path.
			Error(
				errtype.Directory,
				err,
				path)
	}

	if fileInfoWrapper.IsEmptyPath {
		errorWrapper = errnew.Path.Empty()
	}

	return &dirinfo.Result{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             errorWrapper,
		RawPath:           path,
		FileModeRequested: &fileMode,
		IsValidDir:        fileInfoWrapper.IsDirectory,
		HasIssues:         errorWrapper.HasError(),
		IsIgnoredAction:   isIgnoredAction,
		Action:            performingas.CreateAction,
	}
}
