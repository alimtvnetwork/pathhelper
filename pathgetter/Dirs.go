package pathgetter

import (
	"io/ioutil"

	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/pathhelper/normalize"
)

func Dirs(
	separator,
	rootPath string,
	isNormalize bool,
) *errstr.Results {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return errstr.New.Results.ErrorWrapper(errnew.
			Path.
			Error(
				errtype.PathExpand,
				err,
				rootPath2))
	}

	results := make([]string, 0, len(allPaths))

	for _, fileInfo := range allPaths {
		if !fileInfo.IsDir() {
			continue
		}

		combinedPath := rootPath2 +
			separator +
			fileInfo.Name()

		results = append(results, combinedPath)
	}

	return errstr.New.Results.Strings(results)
}
