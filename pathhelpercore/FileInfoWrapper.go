package pathhelpercore

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type FileInfoWrapper struct {
	FileInfo     *os.FileInfo
	Error        errorwrapper.Wrapper
	RawPath      string
	IsDirectory  bool
	IsFile       bool
	IsEmptyPath  bool
	isPathExists *bool
}

func NewFileWrapperInfo(rawPath string) *FileInfoWrapper {
	isEmptyPath := IsEmptyPath(rawPath)

	if isEmptyPath {
		emptyFileError := errorwrapper.NewFilePath(msgtype.InvalidEmptyPathErrorMessage.String(), constants.EmptyString)

		return &FileInfoWrapper{
			FileInfo:    nil,
			Error:       emptyFileError,
			RawPath:     rawPath,
			IsDirectory: false,
			IsFile:      false,
			IsEmptyPath: isEmptyPath,
		}
	}

	fileInfo, error := os.Stat(rawPath)
	isDir := error == nil && fileInfo.IsDir()

	return &FileInfoWrapper{
		FileInfo:    &fileInfo,
		Error:       errorwrapper.NewUsingError(errtype.FileOrDirectoryRelatedExecution, error),
		RawPath:     rawPath,
		IsDirectory: isDir,
		IsFile:      error == nil && !isDir,
		IsEmptyPath: isEmptyPath,
	}
}

func (fileInfoWrapper *FileInfoWrapper) HasError() bool {
	return fileInfoWrapper.Error.HasError()
}

func (fileInfoWrapper *FileInfoWrapper) IsPathExists() bool {
	if nil == fileInfoWrapper.isPathExists {
		isPathExists := !fileInfoWrapper.HasError() && (fileInfoWrapper.IsDirectory || fileInfoWrapper.IsFile)
		fileInfoWrapper.isPathExists = &isPathExists
	}

	return *fileInfoWrapper.isPathExists
}
