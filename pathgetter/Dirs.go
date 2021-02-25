package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func Dirs(
	separator,
	rootPath string,
	isNormalize bool,
) (
	*[]string, *errorwrapper.Wrapper,
) {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return &(constants.EmptyStrings),
			errnew.NewPtr(errtype.FileInfo, err)
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

	return &results, errnew.EmptyPtr
}
