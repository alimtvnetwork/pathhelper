package pathhelper

import (
	"errors"
	"os"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// Create all sub-directories and create the final directory
func CreateDirectoryAll(path string, fileMode os.FileMode) *pathhelpercore.DirectoryResult {
	fileInfoWrapper := GetFileInfoWrapper(path)
	isIgnoredAction := fileInfoWrapper.IsPathExists() || fileInfoWrapper.IsEmptyPath
	var error error

	if !isIgnoredAction {
		error = os.MkdirAll(path, fileMode)
	}

	if fileInfoWrapper.IsEmptyPath {
		error = errors.New(constants.InvalidEmptyPathErrorMessage)
	}

	return &pathhelpercore.DirectoryResult{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             &error,
		RawPath:           path,
		FileModeRequested: fileMode,
		HasIssues:         error != nil,
		IsIgnoredAction:   isIgnoredAction,
		Action:            enums.CreateAction,
	}
}
