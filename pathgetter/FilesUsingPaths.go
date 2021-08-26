package pathgetter

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func FilesUsingPaths(
	isNormalize bool,
	separator string,
	exploringPaths ...string,
) *errstr.Results {
	if exploringPaths == nil {
		return &errstr.Results{
			Values:       &[]string{},
			ErrorWrapper: errnew.EmptyPtr,
		}
	}

	return FilesUsingPathsPtr(
		isNormalize,
		separator,
		&exploringPaths)
}
