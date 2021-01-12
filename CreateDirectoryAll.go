package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// Create all sub-directories and create the final directory
func CreateDirectoryAll(path string, fileMode os.FileMode) *pathhelpercore.DirectoryResult {
	fileInfoWrapper := GetFileInfoWrapper(path)
	isIgnoredAction := fileInfoWrapper.IsPathExists() || fileInfoWrapper.IsEmptyPath
	errorWrapper := errorwrapper.Empty(false)

	if !isIgnoredAction {
		err := os.MkdirAll(path, fileMode)
		errorWrapper = errorwrapper.NewFile(err)
	}

	if fileInfoWrapper.IsEmptyPath {
		errorWrapper = errorwrapper.NewFilePath(msgtype.InvalidEmptyPathErrorMessage.String(), constants.EmptyString)
	}

	return &pathhelpercore.DirectoryResult{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             errorWrapper,
		RawPath:           path,
		FileModeRequested: &fileMode,
		HasIssues:         errorWrapper.HasError(),
		IsIgnoredAction:   isIgnoredAction,
		Action:            enums.CreateAction,
	}
}
