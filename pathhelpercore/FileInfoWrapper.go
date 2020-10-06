package pathhelpercore

import (
	"errors"
	"os"
)

const (
	invalidEmptyPathErrorMessage = "Invalid : Empty path given, cannot process it."
)

type FileInfoWrapper struct {
	FileInfo     *os.FileInfo
	Error        *error
	RawPath      string
	IsDirectory  bool
	IsFile       bool
	IsEmptyPath  bool
	isPathExists *bool
}

func NewFileWrapperInfo(rawPath string) *FileInfoWrapper {
	isEmptyPath := IsEmptyPath(rawPath)

	if isEmptyPath {
		emptyFileError := errors.New(invalidEmptyPathErrorMessage)

		return &FileInfoWrapper{
			FileInfo:    nil,
			Error:       &emptyFileError,
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
		Error:       &error,
		RawPath:     rawPath,
		IsDirectory: isDir,
		IsFile:      error == nil && !isDir,
		IsEmptyPath: isEmptyPath,
	}
}

func (fileInfoWrapper *FileInfoWrapper) HasError() bool {
	return *fileInfoWrapper.Error != nil
}

func (fileInfoWrapper *FileInfoWrapper) IsPathExists() bool {
	if nil == fileInfoWrapper.isPathExists {
		isPathExists := fileInfoWrapper.HasError() && (fileInfoWrapper.IsDirectory || fileInfoWrapper.IsFile)
		fileInfoWrapper.isPathExists = &isPathExists
	}

	return *fileInfoWrapper.isPathExists
}
