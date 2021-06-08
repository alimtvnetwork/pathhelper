package recursiveinternal

import (
	"io/ioutil"
	"os"
	"sync"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/ds"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func GetFilterPaths(
	separator,
	rootPath string,
	isContinueOnEmpty bool,
	filter pathfuncs.Filter,
) (*[]string, *errwrappers.Collection) {
	if rootPath == "" {
		return core.EmptyStringsPtr(), errwrappers.Empty()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	return getFilterPaths(
		separator,
		rootPath,
		fileInfos,
		err,
		isContinueOnEmpty,
		filter)
}

//goland:noinspection GoNilness
func getFilterPaths(
	separator,
	rootPath string,
	initialFileInfos []os.FileInfo,
	err error,
	isContinueOnError bool,
	filter pathfuncs.Filter,
) (*[]string, *errwrappers.Collection) {
	if err != nil && !isContinueOnError {
		errnew.ErrPtr(err).HandleErrorWithRefs(
			msgtype.PathErrorMessage.String(),
			"rootPath",
			rootPath)
	} else if err != nil {
		return core.EmptyStringsPtr(),
			errwrappers.
				Empty().
				AddUsingMessages(
					errtype.FileInfo,
					err.Error())
	}

	if initialFileInfos == nil {
		return core.EmptyStringsPtr(),
			errwrappers.Empty()
	}

	length := len(initialFileInfos)

	if length == 0 {
		return core.EmptyStringsPtr(),
			errwrappers.Empty()
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

		fileName := fileInfo.Name()
		currentPath := rootPath +
			separator +
			fileName

		currentRootFileInfo, err2 := os.Stat(currentPath)
		isContinue := (err2 != nil || currentRootFileInfo == nil) &&
			isContinueOnError

		if isContinue {
			continue
		}

		arg := &pathfuncs.FilterArg{
			RootPath:   currentPath,
			FileName:   fileName,
			FullPath:   currentPath,
			Separator:  separator,
			IsFile:     !currentRootFileInfo.IsDir(),
			IsDirector: currentRootFileInfo.IsDir(),
			FileInfo:   currentRootFileInfo,
		}

		result := filter(arg)

		if result.IsKeep {
			results.
				Paths.
				AddStrings(
					result.FullPath)
		}

		if result.IsBreak {
			return results.
					Paths.ListPtr(),
				results.ErrorWrappersCollector
		}

		fileInfos, err3 := ioutil.ReadDir(currentPath)

		go func() {
			results.AddFilter(
				currentPath,
				fileInfos,
				err3,
				filter)

			wg.Done()
		}()

	}

	wg.Wait()

	list := results.Paths.ListPtr()

	// clearing
	results.Paths = nil
	results = nil

	return list, results.ErrorWrappersCollector
}
