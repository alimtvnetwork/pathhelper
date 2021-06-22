package pathmodifierverify

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathstatlinux"
)

// applyVerifierSinglePathNonRecursive
func applyVerifierSinglePathNonRecursive(
	isExistenceVerify bool,
	isNormalize bool,
	verifier *pathinsfmt.PathVerifier,
	location string,
) *errorwrapper.Wrapper {
	existenceErrorWp := existenceVerifyError(
		isExistenceVerify,
		verifier,
		location)

	if existenceErrorWp.HasError() {
		return existenceErrorWp.ErrorWrapper
	}

	location = normalize.PathUsingSingleIf(
		isNormalize,
		location)

	if verifier.HasRwxInstructions() {
		for _, rwxInstruction := range *verifier.BaseRwxInstructions.RwxInstructions {
			err := chmodhelper.VerifyChmodUsingRwxOwnerGroupOther(
				location,
				&rwxInstruction.RwxOwnerGroupOther)

			if err != nil {
				return errnew.PathMessages(
					errtype.Unexpected,
					location,
					err.Error())
			}
		}
	}

	if verifier.IsGroupNameUserNameBothEmpty() {
		return errnew.EmptyPtr
	}

	pathStat := pathstatlinux.Get(location)

	if pathStat == nil || !pathStat.IsValid {
		//goland:noinspection GoNilness
		return errnew.PathMessages(
			errtype.PathStatFailed,
			location,
			pathStat.ErrorWrapper.FullString())
	}

	verifyUsername := pathStat.User.Name
	if verifier.HasUserName() && !verifier.IsUsername(verifyUsername) {
		return errnew.PathMessages(
			errtype.Unexpected,
			location,
			msgtype.Expecting(
				"Username expectation doesn't meet",
				verifier.UserNameSimple(),
				verifyUsername))
	}

	verifyGroupName := pathStat.Group.Name
	if verifier.HasGroupName() && !verifier.IsGroupName(verifyGroupName) {
		return errnew.PathMessages(
			errtype.Unexpected,
			location,
			msgtype.Expecting(
				"Group expectation doesn't meet",
				verifier.GroupName,
				verifyGroupName))
	}

	return errnew.EmptyPtr
}
