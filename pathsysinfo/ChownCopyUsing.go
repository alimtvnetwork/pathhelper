package pathsysinfo

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/fileinfopath"
)

func ChownCopyUsing(
	srcFullPath *fileinfopath.Instance,
	dstFullPath string,
) *errorwrapper.Wrapper {
	if srcFullPath.IsInvalidPath() {
		return srcFullPath.ErrorWrapper(errtype.ChownUserOrGroupApplyIssue)
	}

	userGroupId := GetPathUserGroupIdUsing(srcFullPath)

	return userGroupId.ApplyChown(dstFullPath)
}
