package pathgetter

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func All(
	isNormalize bool,
	separator string,
	exploringPaths ...string,
) *errstr.Results {
	length := len(exploringPaths)

	if length == 0 {
		return &errstr.Results{
			Values:       &[]string{},
			ErrorWrapper: errnew.EmptyPtr,
		}
	}

	if length == 1 {
		return AllOfSinglePath(
			isNormalize,
			separator,
			(exploringPaths)[0])
	}

	return AllPtr(
		separator,
		isNormalize,
		&exploringPaths)
}
