package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyVerifier(
	isNormalize,
	isRecursiveCheck,
	isSkipCheckingOnInvalid,
	isContinueOnError bool,
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if verifier == nil || len(locations) == 0 {
		return true
	}

	errCount := errCollection.Length()
	locationsNormalized := normalize.PathsUsingSingleIfAsync(
		isNormalize,
		locations)

	locationsWithErrors := recursiveinternal.GetPathsOfPathsIf(
		isRecursiveCheck,
		locationsNormalized,
		isContinueOnError)

	errorCollection2 := locationsWithErrors.ErrorWrappers
	if !isContinueOnError && errorCollection2.HasError() {
		errCollection.AddCollections(errorCollection2)

		return false
	}

	errCollection.AddCollections(errorCollection2)
	existingPathsFileInfoMap := normalize.GetFilterPathsInfoMap(
		false,
		isSkipCheckingOnInvalid,
		*locationsWithErrors.Values)

	applyVerifierInternal(
		isContinueOnError,
		verifier,
		errCollection,
		existingPathsFileInfoMap)

	return errCount == errCollection.Length()
}
