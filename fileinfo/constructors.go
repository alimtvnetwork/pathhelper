package fileinfo

import (
	"io/ioutil"
	"os"
	"path"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

func New(rawPath string) *Wrapper {
	isEmptyPath := ispath.Empty(rawPath)

	if isEmptyPath {
		emptyFileError := errorwrapper.NewFilePath(
			msgtype.InvalidEmptyPathErrorMessage.String(),
			constants.EmptyString)

		return &Wrapper{
			FileInfo:    nil,
			Error:       emptyFileError,
			RawPath:     rawPath,
			IsDirectory: false,
			IsFile:      false,
			IsEmptyPath: isEmptyPath,
		}
	}

	fileInfo, err := os.Stat(rawPath)
	isDir := err == nil && fileInfo.IsDir()

	return &Wrapper{
		FileInfo:    &fileInfo,
		Error:       errorwrapper.NewFile(err),
		RawPath:     rawPath,
		IsDirectory: isDir,
		IsFile:      err == nil && !isDir,
		IsEmptyPath: isEmptyPath,
	}
}

func NewError(
	filePath string,
	err error,
) *Wrapper {
	isFilePathEmpty := ispath.Empty(filePath)

	if err != nil {
		errWrapper := errorwrapper.NewFilePath(
			err.Error(),
			filePath)

		return &Wrapper{
			FileInfo:    nil,
			Error:       errWrapper,
			RawPath:     filePath,
			IsDirectory: false,
			IsFile:      false,
			IsEmptyPath: isFilePathEmpty,
		}
	}

	fileErrWrapper := errorwrapper.NewFilePath(
		errtype.FileOrDirectoryRelatedExecution.String(),
		filePath)

	return &Wrapper{
		FileInfo:    nil,
		Error:       fileErrWrapper,
		RawPath:     filePath,
		IsDirectory: false,
		IsFile:      false,
		IsEmptyPath: isFilePathEmpty,
	}
}

func NewUsingInfo(
	osFileInfo os.FileInfo,
	filePath string,
	err error,
) *Wrapper {
	isFilePathEmpty := ispath.Empty(filePath)

	if err != nil || osFileInfo == nil {
		return NewError(
			filePath,
			err)
	}

	fileErrWrapper := errorwrapper.NewFilePath(
		errtype.FileOrDirectoryRelatedExecution.String(),
		filePath)

	isDir := osFileInfo.IsDir()

	return &Wrapper{
		FileInfo:    &osFileInfo,
		Error:       fileErrWrapper,
		RawPath:     filePath,
		IsDirectory: isDir,
		IsFile:      !isDir,
		IsEmptyPath: isFilePathEmpty,
	}
}

func NewWrappersPtrUsingCapacity(rootPath string, capacity int) *Wrappers {
	collection := make([]Wrapper, 0, capacity)

	return &Wrappers{
		RootPath:            rootPath,
		collection:          &collection,
		directories:         nil,
		files:               nil,
		recursivePaths:      nil,
		Error:               errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}
}

func EmptyWrappers() *Wrappers {
	return NewWrappersPtrUsingCapacity("", 0)
}

func NewWrappersPtr(filePath string) *Wrappers {
	fileInfos, err := ioutil.ReadDir(filePath)

	if err != nil {
		errW := errorwrapper.NewFilePath(
			err.Error(),
			filePath)

		return &Wrappers{
			RootPath:   filePath,
			collection: nil,
			Error:      &errW,
		}
	}

	collection := make([]Wrapper, len(fileInfos))

	for i, fileInfo := range fileInfos {
		newFilePath := path.Join(filePath, fileInfo.Name())
		wrapper := NewUsingInfo(
			fileInfo,
			newFilePath,
			nil)

		collection[i] = *wrapper
	}

	return &Wrappers{
		collection: &collection,
		Error:      errnew.EmptyPtr,
	}
}
