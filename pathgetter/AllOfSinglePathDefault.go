package pathgetter

import (
	"io/ioutil"

	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func AllOfSinglePathDefault(
	isNormalize bool,
	exploringPath string,
) *errstr.Results {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		osconsts.PathSeparator,
		exploringPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		return errstr.New.Results.ErrorWrapper(
			errnew.Path.
				Error(
					errtype.PathExpand,
					err,
					rootPath2))
	}

	results := make([]string, len(allPaths))

	for i, fileInfo := range allPaths {
		results[i] = rootPath2 +
			osconsts.PathSeparator +
			fileInfo.Name()
	}

	return errstr.New.Results.Strings(results)
}
