package namegroup

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func Apply(
	isRecursive bool,
	isContinueOnError bool,
	userNameGroupName *pathinsfmt.UserGroupName,
	paths ...string,
) *errorwrapper.Wrapper {
	if userNameGroupName == nil {
		return nil
	}

	if len(paths) == 0 {
		return nil
	}

	if userNameGroupName.IsUserNameEmpty() {
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
