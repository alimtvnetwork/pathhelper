package pathgetter

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func FilesUsingPaths(
	separator string,
	isNormalize bool,
	exploringPaths ...string,
) *errstr.ResultsWithErrorCollection {
	if exploringPaths == nil {
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	return FilesUsingPathsPtr(
		separator,
		isNormalize,
		&exploringPaths)
}
