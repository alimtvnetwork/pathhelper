package recursiveinternal

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/ds"
)

func GetPaths(
	separator,
	rootPath string,
	isContinueOnEmpty bool,
) (*[]string, *errwrappers.Collection) {
	if rootPath == "" {
		return core.EmptyStringsPtr(), errwrappers.Empty()
	}

	fileInfos, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return &[]string{}, errwrappers.NewCap1().AddPathIssue(
			errtype.FileInfo, err, rootPath)
	}

	return getPaths(
		separator,
		rootPath,
		fileInfos,
		err,
		isContinueOnEmpty)
}

// Probably need refinements
func _getPaths(root string) ([]string, *errwrappers.Collection) {
	paths := make([]string, 0, constants.ArbitraryCapacity64)
	errCollection := errwrappers.NewCap2()
	err := filepath.Walk(
		root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				errCollection.AddError(err)

				return err
			}

			paths = append(paths, path)
			return nil
		},
	)

	errCollection.AddError(err)

	return paths, errCollection
}

func getPaths(
	separator,
	rootPath string,
	initialFileInfos []os.FileInfo,
	err error,
	isContinueOnError bool,
) (*[]string, *errwrappers.Collection) {
	if err != nil && !isContinueOnError {
		return &[]string{},
			errwrappers.NewCap1().
				AddPathIssue(
					errtype.FileInfo,
					err,
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
	results.Paths.AddStrings(rootPath)

	for _, fileInfo := range initialFileInfos {
		// expand and ask for recurse
		if fileInfo == nil && isContinueOnError {
			continue
		}

		//goland:noinspection GoNilness
		currentPath := rootPath +
			separator +
			fileInfo.Name()

		results.
			Paths.
			AddStrings(currentPath)

		//goland:noinspection GoNilness
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

	return list, results.ErrorWrappersCollector
}
