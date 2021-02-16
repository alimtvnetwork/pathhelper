package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/dirinfo"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Create directory and create the final directory
func CreateDirectory(path string, fileMode os.FileMode) *fileinfo.Result {
	fileInfoWrapper := GetFileInfoWrapper(path)
	isIgnoredAction := fileInfoWrapper.IsPathExists() || fileInfoWrapper.IsEmptyPath
	errorWrapper := errnew.Empty

	if !isIgnoredAction {
		err := os.MkdirAll(path, fileMode)
		errorWrapper = errorwrapper.NewFile(err)
	}

	if fileInfoWrapper.IsEmptyPath {
		errorWrapper = errorwrapper.NewFilePath(msgtype.InvalidEmptyPathErrorMessage.String(), constants.EmptyString)
	}

	return &fileinfo.Result{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             errorWrapper,
		RawPath:           path,
		FileModeRequested: &fileMode,
		HasIssues:         errorWrapper.HasError(),
		IsIgnoredAction:   isIgnoredAction,
		Action:            enums.CreateAction,
	}
}
