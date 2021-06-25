package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingFlatPaths(
	isContinueOnError bool,
	verifiers *pathinsfmt.PathVerifiers,
	errCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if verifiers == nil || verifiers.IsEmpty() {
		return true
	}

	existingErrorCount := errCollection.Length()
	locationsNormalized := normalize.PathsUsingSingleIfAsync(
		verifiers.IsNormalize,
		locations)

	locationsWithErrors := recursiveinternal.GetPathsOfPathsIf(
		verifiers.IsRecursiveCheck,
		locationsNormalized,
		isContinueOnError)

	if !isContinueOnError && locationsWithErrors.HasError() {
		errCollection.AddCollections(locationsWithErrors.ErrorWrappers)

		return false
	}

	errCollection.AddCollections(locationsWithErrors.ErrorWrappers)

	existingPathsFileInfoMap := normalize.GetFilterPathsInfoMap(
		false,
		verifiers.IsSkipCheckingOnInvalid,
		*locationsWithErrors.Values)

	// exit immediately
	if !isContinueOnError {
		for _, verifier := range verifiers.PathVerifiers {
			isSuccess = applyVerifierInternal(
				isContinueOnError,
				&verifier,
				errCollection,
				existingPathsFileInfoMap)

			if !isSuccess {
				return false
			}
		}
	}

	// continue on error
	for _, verifier := range verifiers.PathVerifiers {
		applyVerifierInternal(
			isContinueOnError,
			&verifier,
			errCollection,
			existingPathsFileInfoMap)
	}

	isSuccess = existingErrorCount == errCollection.Length()

	return isSuccess
}
