package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingPathVerifiersWithGenericPathsCollection(
	isContinueOnError bool,
	pathVerifiersWithGenericPathsCollection *pathinsfmt.PathVerifiersWithGenericPathsCollection,
) *errorwrapper.Wrapper {
	if pathVerifiersWithGenericPathsCollection == nil ||
		pathVerifiersWithGenericPathsCollection.IsEitherEmpty() {
		return errnew.EmptyPtr
	}

	return ApplyUsingFlatPaths(
		isContinueOnError,
		pathVerifiersWithGenericPathsCollection.PathVerifiers,
		pathVerifiersWithGenericPathsCollection.GenericPathsCollection.LazyFlatPaths())
}
