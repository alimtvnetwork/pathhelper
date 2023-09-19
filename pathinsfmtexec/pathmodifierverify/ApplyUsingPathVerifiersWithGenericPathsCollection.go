package pathmodifierverify

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyUsingPathVerifiersWithGenericPathsCollection(
	isContinueOnError bool,
	pathVerifiersWithGenericPathsCollection *pathinsfmt.PathVerifiersWithGenericPathsCollection,
	errCollection *errwrappers.Collection,
) (isSuccess bool) {
	if pathVerifiersWithGenericPathsCollection == nil ||
		pathVerifiersWithGenericPathsCollection.IsEitherEmpty() {
		return false
	}

	return ApplyUsingFlatPaths(
		isContinueOnError,
		pathVerifiersWithGenericPathsCollection.PathVerifiers,
		errCollection,
		pathVerifiersWithGenericPathsCollection.GenericPathsCollection.LazyFlatPaths())
}
