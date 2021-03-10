package fileinfo

import (
	"io/ioutil"
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/ispath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func New(rawPath, separator string) *Wrapper {
	isEmptyPath := ispath.Empty(rawPath)

	if isEmptyPath {
		emptyFileError := errorwrapper.NewFilePathPtr(
			msgtype.InvalidEmptyPathErrorMessage.String(),
			constants.EmptyString)

		return &Wrapper{
			FileInfo:     nil,
			ErrorWrapper: emptyFileError,
			RawPath:      rawPath,
			IsDirectory:  false,
			IsFile:       false,
			IsEmptyPath:  isEmptyPath,
		}
	}

	fileInfo, err := os.Stat(rawPath)
	isDir := err == nil && fileInfo.IsDir()

	return &Wrapper{
		FileInfo:     &fileInfo,
		ErrorWrapper: errorwrapper.NewFilePtr(err),
		RawPath:      rawPath,
		IsDirectory:  isDir,
		IsFile:       err == nil && !isDir,
		IsEmptyPath:  isEmptyPath,
		Separator:    separator,
	}
}

func NewError(
	filePath, separator string,
	err error,
) *Wrapper {
	isFilePathEmpty := ispath.Empty(filePath)

	if err != nil {
		errWrapper := errorwrapper.NewFilePathPtr(
			err.Error(),
			filePath)

		return &Wrapper{
			FileInfo:     nil,
			ErrorWrapper: errWrapper,
			RawPath:      filePath,
			IsDirectory:  false,
			IsFile:       false,
			IsEmptyPath:  isFilePathEmpty,
			Separator:    separator,
		}
	}

	fileErrWrapper := errorwrapper.NewFilePathPtr(
		errtype.FileOrDirectoryRelatedExecution.String(),
		filePath)

	return &Wrapper{
		FileInfo:     nil,
		ErrorWrapper: fileErrWrapper,
		RawPath:      filePath,
		IsDirectory:  false,
		IsFile:       false,
		IsEmptyPath:  isFilePathEmpty,
		Separator:    separator,
	}
}

func NewUsingInfo(
	osFileInfo os.FileInfo,
	filePath, separator string,
	err error,
) *Wrapper {
	isFilePathEmpty := ispath.Empty(filePath)

	if err != nil || osFileInfo == nil {
		return NewError(
			filePath,
			separator,
			err)
	}

	fileErrWrapper := errorwrapper.NewFilePathPtr(
		errtype.FileOrDirectoryRelatedExecution.String(),
		filePath)

	isDir := osFileInfo.IsDir()

	return &Wrapper{
		FileInfo:     &osFileInfo,
		ErrorWrapper: fileErrWrapper,
		RawPath:      filePath,
		IsDirectory:  isDir,
		IsFile:       !isDir,
		IsEmptyPath:  isFilePathEmpty,
		Separator:    separator,
	}
}

func NewWrappersPtrUsingCapacity(rootPath string, capacity int) *Wrappers {
	collection := make([]*Wrapper, 0, capacity)

	return &Wrappers{
		RootPath:            rootPath,
		Items:               &collection,
		directories:         nil,
		files:               nil,
		recursiveDirs:       nil,
		ErrorWrapper:        errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}
}

func EmptyWrappers() *Wrappers {
	return NewWrappersPtrUsingCapacity("", 0)
}

func NewWrappersPtr(
	filePath, separator string,
	isNormalize bool,
) *Wrappers {
	fileInfos, err := ioutil.ReadDir(filePath)

	if err != nil {
		errW := errorwrapper.NewFilePath(
			err.Error(),
			filePath)

		return &Wrappers{
			RootPath:     filePath,
			Items:        nil,
			ErrorWrapper: &errW,
			Separator:    separator,
		}
	}

	collection := make(
		[]*Wrapper,
		len(fileInfos))

	filePathNormalized := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		filePath,
	)

	for i, fileInfo := range fileInfos {
		newFilePath := filePathNormalized +
			separator +
			fileInfo.Name()

		wrapper := NewUsingInfo(
			fileInfo,
			newFilePath,
			separator,
			nil)

		collection[i] = wrapper
	}

	return &Wrappers{
		Items:        &collection,
		ErrorWrapper: errnew.EmptyPtr,
		Separator:    separator,
	}
}
