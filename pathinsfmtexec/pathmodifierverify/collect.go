package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func collectRecursiveCheckErrors(
	collection *errwrappers.Collection,
	verifier *pathinsfmt.PathVerifier,
	location string,
) {
	recursivePaths, errCollection2 := recursiveinternal.GetPathsWithoutSeparator(
		location,
		false)

	collection.AddCollections(errCollection2)

	if collection.HasError() {
		return
	}

	for _, recursiveLoc := range *recursivePaths {
		errWp := applyVerifierSinglePathNonRecursive(verifier, recursiveLoc)

		collection.AddWrapperPtr(errWp)
	}
}
