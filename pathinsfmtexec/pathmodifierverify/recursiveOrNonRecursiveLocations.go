package pathmodifierverify

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/normalize"
	"gitlab.com/auk-go/pathhelper/recursivepaths"
)

func recursiveOrNonRecursiveLocations(
	isContinueOnError bool,
	isRecursiveCheck bool,
	isNormalize bool,
	locations []string,
) *errstr.Results {
	locationsNormalized := normalize.PathsUsingSingleIfAsync(
		isNormalize,
		locations)

	return recursivepaths.AllOfLocationsIf(
		isRecursiveCheck,
		isContinueOnError,
		locationsNormalized,
	)
}
