package namegroup

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyLinuxOnlyGroup(
	isRecursive bool,
	isContinueOnError bool,
	baseGroupName *pathinsfmt.BaseGroupName,
	paths ...string,
) *errorwrapper.Wrapper {
	if baseGroupName == nil {
		return errnew.EmptyPtr
	}

	if stringutil.IsEmptyOrWhitespace(baseGroupName.GroupName) {
		return errnew.EmptyPtr
	}

	pathsLength := len(paths)

	if pathsLength == 0 {
		return errnew.EmptyPtr
	}

	groupName := baseGroupName.GroupName

	// chgrp groupName path
	cmdPrefix := cmdprefix.ChangeGroup(isRecursive, groupName)

	return applyCmdOnPaths(
		cmdPrefix,
		paths,
		isContinueOnError)
}
