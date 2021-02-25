package pathgetter

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func All(
	separator string,
	isNormalize bool,
	exploringPaths ...string,
) (*[]string, *errorwrapper.Wrapper) {
	length := len(exploringPaths)

	if length == 0 {
		return &(constants.EmptyStrings), errnew.EmptyPtr
	}

	if length == 1 {
		return AllOfSinglePath(
			separator,
			isNormalize,
			(exploringPaths)[0])
	}

	return AllPtr(
		separator,
		isNormalize,
		&exploringPaths)
}
