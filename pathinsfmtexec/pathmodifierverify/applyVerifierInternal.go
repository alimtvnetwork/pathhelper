package pathmodifierverify

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func applyVerifierInternal(
	isContinueOnError bool,
	verifier *pathinsfmt.PathVerifier,
	filteredPathFileInfoMap *chmodhelper.FilteredPathFileInfoMap,
) *errorwrapper.Wrapper {
	if verifier == nil || filteredPathFileInfoMap == nil {
		return errnew.EmptyPtr
	}

	isExistOnError := !isContinueOnError
	errCollection := errwrappers.Empty()

	if filteredPathFileInfoMap.Error != nil && isExistOnError {
		return errnew.NewPtr(
			errtype.PathMissingOrInvalid,
			filteredPathFileInfoMap.Error)
	}

	if filteredPathFileInfoMap.Error != nil && isContinueOnError {
		errCollection.AddTypeError(
			errtype.PathMissingOrInvalid,
			filteredPathFileInfoMap.Error)
	}

	validFileLocations := mics.GetLocationsUsingFilteredPathFileInfoMap(
		filteredPathFileInfoMap)

	if verifier.HasRwxInstructions() {
		executors, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(
			&verifier.BaseRwxInstructions)

		if err != nil && isExistOnError {
			return errnew.NewPtr(errtype.ParsingFailed, err)
		}

		if err != nil && isContinueOnError {
			errCollection.AddTypeError(
				errtype.ParsingFailed,
				err)
		}

		errRwxInstructions := executors.VerifyRwxModifiers(
			isContinueOnError,
			true,
			validFileLocations)

		if errRwxInstructions != nil && isExistOnError {
			return errnew.NewPtr(errtype.ParsingFailed, err)
		}

		if errRwxInstructions != nil && isContinueOnError {
			errCollection.AddTypeError(
				errtype.ParsingFailed,
				err)
		}
	}

	if isExistOnError {
		for _, location := range validFileLocations {
			errWp := applyVerifierSinglePathNonRecursiveUserGroupVerify(
				verifier,
				location)

			if errWp.HasError() {
				return errWp
			}
		}
	}

	for _, location := range validFileLocations {
		errWp := applyVerifierSinglePathNonRecursiveUserGroupVerify(
			verifier,
			location)

		if errWp.HasError() {
			errCollection.AddWrapperPtr(errWp)
		}
	}

	return errCollection.GetAsErrorWrapperPtr()
}
