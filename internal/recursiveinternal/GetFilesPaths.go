package recursiveinternal

import (
	"io/ioutil"
	"os"
	"sync"

	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/internal/ds"
)

func GetFilesPaths(
	separator,
	rootPath string,
	isContinueOnEmpty bool,
) *errstr.ResultsWithErrorCollection {
	if rootPath == "" {
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	return getFilesPaths(
		separator,
		rootPath,
		fileInfos,
		err,
		isContinueOnEmpty)
}

//goland:noinspection GoNilness
func getFilesPaths(
	separator,
	rootPath string,
	initialFileInfos []os.FileInfo,
	err error,
	isContinueOnError bool,
) *errstr.ResultsWithErrorCollection {
	if err != nil && !isContinueOnError {
		errnew.ErrPtr(err).HandleErrorWithRefs(
			msgtype.FileErrorMessage.String(),
			"rootPath",
			rootPath)
	} else if err != nil {
		return errstr.
			NewResultsWithErrorCollectionUsingTypeErrorPtr(
				errtype.FileInfo, err)
	}

	if initialFileInfos == nil {
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	length := len(initialFileInfos)

	if length == 0 {
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	nestedPathsParam := ds.NestedPathsParam{
		IsContinueOnError: isContinueOnError,
		Separator:         separator,
	}

	results := nestedPathsParam.Results()

	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, fileInfo := range initialFileInfos {
		// expand and ask for recurse
		if fileInfo == nil && isContinueOnError {
			continue
		}

		currentPath := rootPath +
			separator +
			fileInfo.Name()

		if !fileInfo.IsDir() {
			results.
				Paths.
				AddStrings(currentPath)
		}

		if fileInfo.IsDir() {
			fileInfos, err2 := ioutil.
				ReadDir(currentPath)

			go func() {
				results.AddFile(
					currentPath,
					fileInfos,
					err2)

				wg.Done()
			}()
		} else {
			wg.Done()
		}
	}

	wg.Wait()

	return results.
		ToResultsWithErrorCollection()
}
