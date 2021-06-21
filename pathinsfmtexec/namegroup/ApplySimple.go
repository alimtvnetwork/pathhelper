package namegroup

import (
	"gitlab.com/evatix-go/core/typesconv"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplySimple(
	isRecursive bool,
	isContinueOnError bool,
	userName *string,
	groupName *string,
	paths ...string,
) *errorwrapper.Wrapper {
	if userName == nil && groupName == nil {
		return errnew.EmptyPtr
	}

	if len(paths) == 0 {
		return errnew.EmptyPtr
	}

	groupNameSimple := typesconv.StringPtrToSimple(groupName)

	userNameGroup := pathinsfmt.BaseUserNamePlusGroupName{
		BaseGroupName: pathinsfmt.BaseGroupName{GroupName: groupNameSimple},
		UserName:      userName,
	}

	return Apply(
		isRecursive,
		isContinueOnError,
		&userNameGroup,
		paths...)
}
