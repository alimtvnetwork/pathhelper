package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
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
	locations []string,
) *errorwrapper.Wrapper {
	if verifier == nil || len(locations) == 0 {
		return errnew.EmptyPtr
	}

	locationsNormalized := normalize.PathsUsingSingleIfAsync(
		isNormalize,
		locations)

	locationsErrors := recursiveinternal.GetPathsOfPathsIf(
		isRecursiveCheck,
		locationsNormalized,
		isContinueOnError)

	if !isContinueOnError && locationsErrors.HasError() {
		return locationsErrors.
			ErrorWrappers.
			GetAsErrorWrapperPtr()
	}

	existingPathsFileInfoMap := normalize.GetFilterPathsInfoMap(
		false,
		isSkipCheckingOnInvalid,
		*locationsErrors.Values)

	errWpFinal := applyVerifierInternal(
		isContinueOnError,
		verifier,
		existingPathsFileInfoMap)

	return locationsErrors.
		ErrorWrappers.
		AddWrapperPtr(errWpFinal).
		GetAsErrorWrapperPtr()
}
