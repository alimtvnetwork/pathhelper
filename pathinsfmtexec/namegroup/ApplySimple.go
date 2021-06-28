package namegroup

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplySimple(
	isRecursive bool,
	isContinueOnError bool,
	userName string,
	groupName string,
	paths ...string,
) *errorwrapper.Wrapper {
	if userName == "" && groupName == "" {
		return errnew.EmptyPtr
	}

	if len(paths) == 0 {
		return errnew.EmptyPtr
	}

	userNameGroup := pathinsfmt.UserGroupName{
		BaseGroupName: pathinsfmt.BaseGroupName{GroupName: groupName},
		UserName:      userName,
	}

	return Apply(
		isRecursive,
		isContinueOnError,
		&userNameGroup,
		paths...)
}
