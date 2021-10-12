package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/dirinfo"
	"gitlab.com/evatix-go/pathhelper/performingas"
)

// AllRecurse Create all sub-directories and create the final directory
func AllRecurse(path string, fileMode os.FileMode) *dirinfo.Result {
	fileInfoWrapper := pathhelper.GetFileInfoWrapper(path)
	isIgnoredAction := fileInfoWrapper.IsPathExists() || fileInfoWrapper.IsEmptyPath
	errorWrapper := errnew.EmptyPtr

	if !isIgnoredAction {
		err := os.MkdirAll(path, fileMode)
		errorWrapper = errnew.Path(errtype.Directory, err, path)
	}

	if fileInfoWrapper.IsEmptyPath {
		errorWrapper = errnew.EmptyFilePath
	}

	return &dirinfo.Result{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             errorWrapper,
		RawPath:           path,
		IsValidDir:        fileInfoWrapper.IsDirectory,
		FileModeRequested: &fileMode,
		HasIssues:         errorWrapper.HasError(),
		IsIgnoredAction:   isIgnoredAction,
		Action:            performingas.CreateAction,
	}
}
