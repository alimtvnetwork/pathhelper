package createdir

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/dirinfo"
	"gitlab.com/evatix-go/pathhelper/performing"
)

// Create all sub-directories and create the final directory
func AllRecurse(path string, fileMode os.FileMode) *dirinfo.Result {
	fileInfoWrapper := pathhelper.GetFileInfoWrapper(path)
	isIgnoredAction := fileInfoWrapper.IsPathExists() || fileInfoWrapper.IsEmptyPath
	errorWrapper := errnew.Empty

	if !isIgnoredAction {
		err := os.MkdirAll(path, fileMode)
		errorWrapper = errorwrapper.NewFile(err)
	}

	if fileInfoWrapper.IsEmptyPath {
		errorWrapper = errorwrapper.NewFilePath(msgtype.InvalidEmptyPathErrorMessage.String(), constants.EmptyString)
	}

	return &dirinfo.Result{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             &errorWrapper,
		RawPath:           path,
		IsValidDir:        fileInfoWrapper.IsDirectory,
		FileModeRequested: &fileMode,
		HasIssues:         errorWrapper.HasError(),
		IsIgnoredAction:   isIgnoredAction,
		Action:            performing.CreateAction,
	}
}
