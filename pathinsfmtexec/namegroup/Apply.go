package namegroup

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(
	isRecursive bool,
	isContinueOnError bool,
	userNameGroupName *pathinsfmt.BaseUserNamePlusGroupName,
	paths ...string,
) *errorwrapper.Wrapper {
	if userNameGroupName == nil {
		return errnew.EmptyPtr
	}

	if len(paths) == 0 {
		return errnew.EmptyPtr
	}

	if stringutil.IsEmptyOrWhitespacePtr(userNameGroupName.UserName) {
		return ApplyLinuxOnlyGroup(
			isRecursive,
			isContinueOnError,
			&userNameGroupName.BaseGroupName,
			paths...)
	}

	return applyLinuxUserGroupBoth(
		isRecursive,
		isContinueOnError,
		userNameGroupName,
		paths...)
}
