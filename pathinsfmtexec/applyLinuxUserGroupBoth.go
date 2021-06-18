package pathinsfmtexec

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
	"gitlab.com/evatix-go/pathhelper/internal/deferrwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func applyLinuxUserGroupBoth(
	isRecursive bool,
	isContinueOnError bool,
	userNameGroupName *pathinsfmt.BaseUserNamePlusGroupName,
	paths ...string,
) *errorwrapper.Wrapper {
	if userNameGroupName == nil {
		return errnew.EmptyPtr
	}

	if stringutil.IsEmptyOrWhitespace(userNameGroupName.GroupName) {
		return deferrwrappers.CannotApplyChmodWithSingleParameter
	}

	if stringutil.IsEmptyOrWhitespacePtr(userNameGroupName.UserName) {
		return deferrwrappers.CannotApplyChmodWithSingleParameter
	}

	pathsLength := len(paths)

	if pathsLength == 0 {
		return errnew.EmptyPtr
	}

	groupName := userNameGroupName.GroupName
	userName := *userNameGroupName.UserName

	// chown -R $user:$group /dir
	cmdPrefix := cmdprefix.ChownUser(
		isRecursive,
		userName,
		groupName)

	return applyCmdOnPaths(
		cmdPrefix,
		paths,
		isContinueOnError)
}
