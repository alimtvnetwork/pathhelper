package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func AllOfSinglePath(
	separator string,
	isNormalize bool,
	exploringPath string,
) (*[]string, *errorwrapper.Wrapper) {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		exploringPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return &(constants.EmptyStrings),
			errnew.NewPtr(errtype.FileInfo, err)
	}

	results := make([]string, len(allPaths))

	for i, fileInfo := range allPaths {
		results[i] = rootPath2 +
			separator +
			fileInfo.Name()
	}

	return &results, errnew.EmptyPtr
}
