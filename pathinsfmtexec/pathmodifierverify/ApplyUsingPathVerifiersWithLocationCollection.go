package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingPathVerifiersWithLocationCollection(
	isContinueOnError bool,
	verifiersWithLocations *pathinsfmt.PathVerifiersWithLocationCollection,
) *errorwrapper.Wrapper {
	if verifiersWithLocations == nil ||
		verifiersWithLocations.IsEitherEmpty() {
		return errnew.EmptyPtr
	}

	return ApplyUsingFlatPaths(
		isContinueOnError,
		verifiersWithLocations.PathVerifiers,
		verifiersWithLocations.LocationCollection.LazyFlatPaths())
}
