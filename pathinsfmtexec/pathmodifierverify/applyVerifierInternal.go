package pathmodifierverify

import (
	"strings"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/internal/mics"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func applyVerifierInternal(
	isContinueOnError bool,
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	filteredPathFileInfoMap *chmodhelper.FilteredPathFileInfoMap,
) (isSuccess bool) {
	if verifier == nil || filteredPathFileInfoMap == nil {
		return true
	}

	if isContinueOnError {
		return applyVerifierContinueOnErrorInternal(
			verifier,
			errCollection,
			filteredPathFileInfoMap)
	}

	stateTracker := errCollection.StateTracker()
	if filteredPathFileInfoMap.Error != nil {
		errCollection.AddTypeError(
			errtype.PathMissingOrInvalid,
			filteredPathFileInfoMap.Error)

		return false
	}

	validFileLocations := mics.GetLocationsUsingFilteredPathFileInfoMap(
		filteredPathFileInfoMap)

	// Rwx verify
	if verifier.HasRwxInstructions() {
		executors, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(
			&verifier.BaseRwxInstructions)

		if err != nil {
			errCollection.AddTypeError(
				errtype.ParsingFailed, err)

			return false
		}

		errRwxInstructions := executors.VerifyRwxModifiers(
			isContinueOnError,
			true,
			validFileLocations)

		if errRwxInstructions != nil {
			errCollection.AddTypeError(
				errtype.RwxMismatch,
				errRwxInstructions)

			return false
		}
	}

	// immediately exit on error, user+groups verify.
	hasAnyUserGroupValidation := len(validFileLocations) > 0 &&
		verifier.UserGroupName.HasUserNameOrGroup()

	if hasAnyUserGroupValidation && osconsts.IsWindows {
		errCollection.AddUsingMessages(
			errtype.NotSupportInWindows,
			errtype.ChownUserOrGroupApplyIssue.String(),
			"Cannot verify valid locations:",
			strings.Join(validFileLocations, constants.CommaSpace))

		return false
	}

	// immediately exit on error, user+groups verify.
	for _, location := range validFileLocations {
		isSuccess = applyVerifierSinglePathNonRecursiveUserGroupVerify(
			verifier,
			errCollection,
			location)

		if !isSuccess {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
