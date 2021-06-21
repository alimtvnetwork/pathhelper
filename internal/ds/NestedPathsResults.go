package ds

import (
	"io/ioutil"
	"os"
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

const (
	FilePath = " ,\n\tfile path :"
)

type NestedPathsResults struct {
	param                  *NestedPathsParam
	Paths                  *corestr.LinkedCollections
	ErrorWrappersCollector *errwrappers.Collection
}

func NewNestedPathsResults(param *NestedPathsParam) *NestedPathsResults {
	paths := corestr.NewLinkedCollections()

	return &NestedPathsResults{
		param:                  param,
		Paths:                  paths,
		ErrorWrappersCollector: errwrappers.Empty(),
	}
}

func (nestedPathsResults *NestedPathsResults) ToResultsWithErrorCollection() *errstr.ResultsWithErrorCollection {
	return &errstr.ResultsWithErrorCollection{
		Values:        nestedPathsResults.Paths.ListPtr(),
		ErrorWrappers: nestedPathsResults.ErrorWrappersCollector,
	}
}

func (nestedPathsResults *NestedPathsResults) Add(
	currentRootPath string,
	infos []os.FileInfo,
	err error,
) {
	hasError := err != nil
	isErrContinue := nestedPathsResults.param.IsContinueOnError
	errWrappers := nestedPathsResults.ErrorWrappersCollector
	if hasError {
		nestedPathsResults.ErrorWrappersCollector.AddPathIssue(
			errtype.File,
			err,
			FilePath+
				currentRootPath)
	}

	if (isErrContinue && hasError) || (infos == nil && !hasError) {
		return
	}

	nestedFileInfosLength := len(infos)
	if nestedFileInfosLength == 0 {
		return
	}

	collection := corestr.NewCollection(nestedFileInfosLength)

	wg2 := &sync.WaitGroup{}
	wg2.Add(nestedFileInfosLength)

	for _, fileInfo := range infos {
		currentPath := currentRootPath +
			nestedPathsResults.param.Separator +
			fileInfo.Name()

		collection.
			Add(currentPath)

		// expand and ask for recurse
		isEmptyPath := currentPath == ""
		if isEmptyPath {
			errWrappers.AddUsingMsg(
				errtype.UnexpectedFilePath,
				consts.FilePathEmpty)
		}

		if isEmptyPath && isErrContinue {
			continue
		} else if isEmptyPath && !isErrContinue {
			return
		}

		if !isDir(currentPath){
			wg2.Done()

			nestedPathsResults.Paths.AddLock(collection)

			return
		}

		fileInfos, err2 := ioutil.ReadDir(currentPath)

		go func() {
			nestedPathsResults.Add(
				currentPath,
				fileInfos,
				err2,
			)

			wg2.Done()
		}()
	}

	wg2.Wait()
	nestedPathsResults.Paths.AddLock(collection)
}

func (nestedPathsResults *NestedPathsResults) AddFilter(
	currentRootPath string,
	infos []os.FileInfo,
	err error,
	filter pathfuncs.Filter,
) {
	hasError := err != nil
	isErrContinue := nestedPathsResults.param.IsContinueOnError
	errWrappers := nestedPathsResults.ErrorWrappersCollector
	if hasError {
		nestedPathsResults.ErrorWrappersCollector.AddUsingMessages(
			errtype.File,
			err.Error(),
			FilePath,
			currentRootPath)
	}

	if (isErrContinue && hasError) ||
		(infos == nil && !hasError) {
		return
	}

	nestedFileInfosLength := len(infos)
	if nestedFileInfosLength == 0 {
		return
	}

	collection := corestr.NewCollection(
		nestedFileInfosLength)

	wg2 := &sync.WaitGroup{}
	wg2.Add(nestedFileInfosLength)

	for _, fileInfo := range infos {
		fileName := fileInfo.Name()
		currentPath := currentRootPath +
			nestedPathsResults.param.Separator +
			fileName
		arg := &pathfuncs.FilterArg{
			RootPath:   currentRootPath,
			FileName:   fileName,
			FullPath:   currentPath,
			Separator:  nestedPathsResults.param.Separator,
			IsFile:     !fileInfo.IsDir(),
			IsDirector: fileInfo.IsDir(),
			FileInfo:   fileInfo,
		}

		result := filter(arg)

		if result.IsKeep {
			collection.
				Add(currentPath)
		}

		if result.IsBreak {
			return
		}

		// expand and ask for recurse
		isEmptyPath := currentPath == ""
		if isEmptyPath {
			errWrappers.AddUsingMsg(
				errtype.UnexpectedFilePath,
				consts.FilePathEmpty)
		}

		if isEmptyPath && isErrContinue {
			continue
		} else if isEmptyPath && !isErrContinue {
			return
		}

		if fileInfo.IsDir() {
			fileInfos, err2 := ioutil.ReadDir(currentPath)

			go func() {
				nestedPathsResults.AddFilter(
					currentPath,
					fileInfos,
					err2,
					filter,
				)

				wg2.Done()
			}()
		} else {
			wg2.Done()
		}
	}

	wg2.Wait()
	nestedPathsResults.Paths.AddLock(collection)
}

func (nestedPathsResults *NestedPathsResults) AddFile(
	currentRootPath string,
	infos []os.FileInfo,
	err error,
) {
	hasError := err != nil
	isErrContinue := nestedPathsResults.param.IsContinueOnError
	errWrappers := nestedPathsResults.ErrorWrappersCollector
	if hasError {
		nestedPathsResults.ErrorWrappersCollector.AddUsingMessages(
			errtype.File,
			err.Error(),
			FilePath,
			currentRootPath)
	}

	if (isErrContinue && hasError) ||
		(infos == nil && !hasError) {
		return
	}

	nestedFileInfosLength := len(infos)
	if nestedFileInfosLength == 0 {
		return
	}

	collection := corestr.NewCollection(
		nestedFileInfosLength)

	wg2 := &sync.WaitGroup{}
	wg2.Add(nestedFileInfosLength)

	for _, fileInfo := range infos {
		currentPath := currentRootPath +
			nestedPathsResults.param.Separator +
			fileInfo.Name()

		if !fileInfo.IsDir() {
			collection.
				Add(currentPath)
		}

		// expand and ask for recurse
		isEmptyPath := currentPath == ""
		if isEmptyPath {
			errWrappers.AddUsingMsg(
				errtype.UnexpectedFilePath,
				consts.FilePathEmpty)
		}

		if isEmptyPath && isErrContinue {
			continue
		} else if isEmptyPath && !isErrContinue {
			return
		}

		if fileInfo.IsDir() {
			fileInfos, err2 := ioutil.ReadDir(currentPath)

			go func() {
				nestedPathsResults.AddFile(
					currentPath,
					fileInfos,
					err2,
				)

				wg2.Done()
			}()
		} else {
			wg2.Done()
		}
	}

	wg2.Wait()
	nestedPathsResults.Paths.AddLock(collection)
}

func (nestedPathsResults *NestedPathsResults) AddDirectory(
	currentRootPath string,
	infos []os.FileInfo,
	err error,
) {
	hasError := err != nil
	isErrContinue := nestedPathsResults.param.IsContinueOnError
	errWrappers := nestedPathsResults.ErrorWrappersCollector
	if hasError {
		nestedPathsResults.ErrorWrappersCollector.AddUsingMessages(
			errtype.File,
			err.Error(),
			FilePath,
			currentRootPath)
	}

	if (isErrContinue && hasError) || (infos == nil && !hasError) {
		return
	}

	nestedFileInfosLength := len(infos)
	if nestedFileInfosLength == 0 {
		return
	}

	collection := corestr.NewCollection(nestedFileInfosLength)

	wg2 := &sync.WaitGroup{}
	wg2.Add(nestedFileInfosLength)

	for _, fileInfo := range infos {
		currentPath := currentRootPath +
			nestedPathsResults.param.Separator +
			fileInfo.Name()

		if fileInfo.IsDir() {
			collection.
				Add(currentPath)
		}

		// expand and ask for recurse
		isEmptyPath := currentPath == ""
		if isEmptyPath {
			errWrappers.AddUsingMsg(
				errtype.UnexpectedFilePath,
				consts.FilePathEmpty)
		}

		if isEmptyPath && isErrContinue {
			continue
		} else if isEmptyPath && !isErrContinue {
			return
		}

		fileInfos, err2 := ioutil.ReadDir(currentPath)

		go func() {
			nestedPathsResults.AddDirectory(
				currentPath,
				fileInfos,
				err2,
			)

			wg2.Done()
		}()
	}

	wg2.Wait()
	nestedPathsResults.Paths.AddLock(collection)
}
