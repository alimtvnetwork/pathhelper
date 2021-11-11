package pathgetter

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func All(
	isNormalize bool,
	separator string,
	exploringPaths ...string,
) *errstr.Results {
	length := len(exploringPaths)

	if length == 0 {
		return errstr.Empty.Results()
	}

	if length == 1 {
		return AllOfSinglePath(
			isNormalize,
			separator,
			exploringPaths[0])
	}

	return AllPtr(
		separator,
		isNormalize,
		exploringPaths)
}
