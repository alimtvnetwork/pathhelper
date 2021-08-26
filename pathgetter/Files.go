package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Files(
	isNormalize bool,
	separator,
	rootPath string,
) *errstr.Results {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return &errstr.Results{
			Values: &[]string{},
			ErrorWrapper: errnew.Path(
				errtype.PathExpand,
				err,
				rootPath2),
		}
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

	return &errstr.Results{
		Values:       &results,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
