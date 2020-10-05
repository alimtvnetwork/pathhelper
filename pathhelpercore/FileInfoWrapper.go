package pathhelpercore

import (
	"os"
)

type FileInfoWrapper struct {
	FileInfo    *os.FileInfo
	Error       *error
	RawPath     string
	IsDirectory bool
	IsFile      bool
	IsEmptyPath bool
}

func NewFileWrapperInfo(rawPath string) *FileInfoWrapper {
	isEmptyPath := IsEmptyPath(rawPath)

	if isEmptyPath {
		panic("Given path is empty.")
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
	return *fileInfoWrapper.Error == nil && (fileInfoWrapper.IsDirectory || fileInfoWrapper.IsFile)
}
