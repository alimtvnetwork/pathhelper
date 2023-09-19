package pathmodifierverify

import (
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathstatlinux"
)

// applyVerifierSinglePathNonRecursiveUserGroupVerify
func applyVerifierSinglePathNonRecursiveUserGroupVerify(
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	location string,
) (isSuccess bool) {
	if verifier.IsGroupNameUserNameBothEmpty() {
		return true
	}

	errCount := errCollection.Length()
	pathStat := pathstatlinux.Get(location)

	if pathStat == nil || !pathStat.IsValidParsing {
		//goland:noinspection GoNilness
		errCollection.AddWrapperPtr(
			pathStat.ErrorWrapper)

		return false
	}

	verifyUsername := pathStat.User.Name
	if verifier.HasUserName() && !verifier.IsUsername(verifyUsername) {
		errCollection.AddPathIssueMessages(
			errtype.Unexpected,
			location,
			errcore.ExpectingSimpleNoType(
				"Username expectation doesn't meet",
				verifier.UserNameSimple(),
				verifyUsername))
	}

	verifyGroupName := pathStat.Group.Name
	if verifier.HasGroupName() && !verifier.IsGroupName(verifyGroupName) {
		errCollection.AddPathIssueMessages(
			errtype.Unexpected,
			location,
			errcore.ExpectingSimpleNoType(
				"Group expectation doesn't meet",
				verifier.GroupName,
				verifyGroupName))
	}

	return errCount == errCollection.Length()
}
