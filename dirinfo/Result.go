package dirinfo

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
	"gitlab.com/evatix-go/pathhelper/performing"
)

type Result struct {
	FileInfoWrapper   *fileinfo.Wrapper
	Error             *errorwrapper.Wrapper
	RawPath           string
	IsValidDir        bool
	FileModeRequested *os.FileMode
	HasIssues         bool
	IsIgnoredAction   bool
	Action            performing.Action
}

func Empty() *Result {
	return EmptyUsingInfo(nil)
}

func EmptyUsingInfo(fileWrapperInfo *fileinfo.Wrapper) *Result {
	return &Result{
		FileInfoWrapper:   fileWrapperInfo,
		Error:             errnew.EmptyPtr,
		IsValidDir:        false,
		RawPath:           "",
		FileModeRequested: nil,
		HasIssues:         false,
		IsIgnoredAction:   true,
		Action:            performing.EmptyDirectoryResult,
	}
}

func New(filePath string) *Result {
	isFilePathEmpty := filePath == ""
	fileInfo, err := os.Stat(filePath)
	errWrapper := errorwrapper.NewDirectory(err)
	isErrorEmpty := errWrapper.IsEmpty()

	fileInfoWrapper := &fileinfo.Wrapper{
		FileInfo:    &fileInfo,
		Error:       errWrapper,
		RawPath:     filePath,
		IsDirectory: isErrorEmpty && fileInfo.IsDir(),
		IsFile:      isErrorEmpty && !fileInfo.IsDir(),
		IsEmptyPath: isFilePathEmpty,
	}

	if !fileInfoWrapper.IsDirectory {
		return &Result{
			FileInfoWrapper:   fileInfoWrapper,
			Error:             &errWrapper,
			RawPath:           filePath,
			FileModeRequested: nil,
			IsValidDir:        false,
			HasIssues:         !isErrorEmpty,
			IsIgnoredAction:   true,
			Action:            performing.EmptyDirectoryResult,
		}
	}

	fileMode := fileInfo.Mode()

	return &Result{
		FileInfoWrapper:   fileInfoWrapper,
		Error:             &errWrapper,
		RawPath:           filePath,
		IsValidDir:        true,
		FileModeRequested: &fileMode,
		HasIssues:         !isErrorEmpty,
		IsIgnoredAction:   true,
		Action:            performing.NoAction,
	}
}

func (receiver *Result) IsEmpty() bool {
	return receiver == nil ||
		receiver.Action.IsEmptyDirectoryResult()
}

func (receiver *Result) HasValidDir() bool {
	return receiver != nil &&
		receiver.IsValidDir &&
		!receiver.HasIssues
}
