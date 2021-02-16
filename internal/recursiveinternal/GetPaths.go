package recursiveinternal

import (
	"io/ioutil"
	"os"
	"sync"

	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/internal/ds"
)

func GetPaths(
	separator,
	rootPath string,
	isContinueOnEmpty bool,
) (*[]string, *errwrappers.Collection) {
	if rootPath == "" {
		return consts.EmptyStringsResultPtr(), errwrappers.Empty()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	return getPaths(
		separator,
		rootPath,
		fileInfos,
		err,
		isContinueOnEmpty)
}

func getPaths(
	separator,
	rootPath string,
	initialFileInfos []os.FileInfo,
	err error,
	isContinueOnError bool,
) (*[]string, *errwrappers.Collection) {
	if err != nil && !isContinueOnError {
		errnew.ErrPtr(err).HandleErrorWithRefs(
			msgtype.FileErrorMessage.String(),
			"rootPath",
			rootPath)
	} else if err != nil {
		return consts.EmptyStringsResultPtr(),
			errwrappers.
				Empty().
				AddUsingMessages(
					errtype.FileInfo,
					err.Error())
	}

	if initialFileInfos == nil {
		return consts.EmptyStringsResultPtr(),
			errwrappers.Empty()
	}

	length := len(initialFileInfos)

	if length == 0 {
		return consts.EmptyStringsResultPtr(),
			errwrappers.Empty()
	}

	nestedPathsParam := ds.NestedPathsParam{
		IsContinueOnError: isContinueOnError,
		Separator:         separator,
	}

	results := nestedPathsParam.Results()

	wg := &sync.WaitGroup{}
	wg.Add(length)
	results.Paths.AddStrings(rootPath)

	for _, fileInfo := range initialFileInfos {
		// expand and ask for recurse
		if fileInfo == nil && isContinueOnError {
			continue
		}

		currentPath := rootPath +
			separator +
			fileInfo.Name()

		results.
			Paths.
			AddStrings(currentPath)

		if fileInfo.IsDir() {
			fileInfos, err2 := ioutil.
				ReadDir(currentPath)

			go func() {
				results.Add(
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

	list := results.Paths.ListPtr()

	// clearing
	results.Paths = nil

	return list, results.ErrWrappers
}
