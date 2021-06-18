package pathinsfmtexec

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func applyLinuxOnlyGroup(
	isRecursive bool,
	isContinueOnError bool,
	userNameGroupName *pathinsfmt.BaseUserNamePlusGroupName,
	paths ...string,
) *errorwrapper.Wrapper {
	if userNameGroupName == nil {
		return errnew.EmptyPtr
	}

	if stringutil.IsEmptyOrWhitespace(userNameGroupName.GroupName) {
		return errnew.EmptyPtr
	}

	pathsLength := len(paths)

	if pathsLength == 0 {
		return errnew.EmptyPtr
	}

	groupName := userNameGroupName.GroupName

	// chgrp groupName path
	cmdPrefix := cmdprefix.ChangeGroup(isRecursive, groupName)

	return applyCmdOnPaths(
		cmdPrefix,
		paths,
		isContinueOnError)
}
