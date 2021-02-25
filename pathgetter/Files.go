package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Files(
	separator,
	rootPath string,
	isNormalize bool,
) *errstr.ResultsWithErrorCollection {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return errstr.
			NewResultsWithErrorCollectionUsingTypePtr(
				errtype.FileInfo)
	}

	results := make([]string, 0, len(allPaths))

	for _, fileInfo := range allPaths {
		if fileInfo.IsDir() {
			continue
		}

		combinedPath := rootPath2 +
			separator +
			fileInfo.Name()

		results = append(results, combinedPath)
	}

	return &errstr.ResultsWithErrorCollection{
		Values:        &results,
		ErrorWrappers: errwrappers.Empty(),
	}
}
