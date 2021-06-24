package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingFlatPaths(
	isContinueOnError bool,
	verifiers *pathinsfmt.PathVerifiers,
	locations []string,
) *errorwrapper.Wrapper {
	if verifiers == nil || verifiers.IsEmpty() {
		return errnew.EmptyPtr
	}

	locationsNormalized := normalize.PathsUsingSingleIfAsync(
		verifiers.IsNormalize,
		locations)

	locationsErrors := recursiveinternal.GetPathsOfPathsIf(
		verifiers.IsRecursiveCheck,
		locationsNormalized,
		isContinueOnError)

	if !isContinueOnError && locationsErrors.HasError() {
		return locationsErrors.
			ErrorWrappers.
			GetAsErrorWrapperPtr()
	}

	existingPathsFileInfoMap := normalize.GetFilterPathsInfoMap(
		false,
		verifiers.IsSkipCheckingOnInvalid,
		*locationsErrors.Values)

	if !isContinueOnError {
		for _, verifier := range verifiers.PathVerifiers {
			errWp := applyVerifierInternal(
				isContinueOnError,
				&verifier,
				existingPathsFileInfoMap)

			locationsErrors.ErrorWrappers.AddWrapperPtr(errWp)
		}

		return locationsErrors.ErrorWrappers.GetAsErrorWrapperPtr()
	}

	for _, verifier := range verifiers.PathVerifiers {
		errWp := applyVerifierInternal(
			isContinueOnError,
			&verifier,
			existingPathsFileInfoMap)

		if errWp.HasError() {
			return errWp
		}
	}

	return errnew.EmptyPtr
}
