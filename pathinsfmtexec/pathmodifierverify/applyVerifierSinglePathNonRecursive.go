package pathmodifierverify

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathstatlinux"
)

// applyVerifierSinglePathNonRecursive
func applyVerifierSinglePathNonRecursive(
	verifier *pathinsfmt.PathVerifier,
	location string,
) *errorwrapper.Wrapper {
	if verifier == nil || len(location) == 0 {
		return errnew.EmptyPtr
	}

	workingPath := normalize.PathUsingSingleIf(
		verifier.IsNormalize,
		location)

	isFileExist := fsinternal.IsPathExists(workingPath)
	isFileMissing := !isFileExist

	if verifier.IsSkipCheckingOnNonExist && isFileMissing {
		return errnew.EmptyPtr
	}

	if !verifier.IsSkipCheckingOnNonExist && isFileMissing {
		return errnew.PathMessages(
			errtype.PathNotFound,
			workingPath,
			"Use IsSkipCheckingOnNonExist to true skip the error.")
	}

	if verifier.HasRwxInstructions() {
		for _, rwxInstruction := range *verifier.BaseRwxInstructions.RwxInstructions {
			err := chmodhelper.VerifyChmodUsingRwxOwnerGroupOther(
				location,
				&rwxInstruction.RwxOwnerGroupOther)

			if err != nil {
				return errnew.PathMessages(
					errtype.Unexpected,
					location,
					err.Error(),
					"Chmod Mismatch, expecting: ",
					rwxInstruction.RwxOwnerGroupOther.String())
			}
		}
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
	if !verifier.IsUsername(verifyUsername) {
		return errnew.PathMessages(
			errtype.Unexpected,
			location,
			msgtype.Expecting(
				"Username expectation doesn't meet",
				verifier.UserNameSimple(),
				verifyUsername))
	}

	verifyGroupName := pathStat.Group.Name
	if !verifier.IsGroupName(verifyGroupName) {
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
