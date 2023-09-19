package pathsysinfo

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/fileinfopath"
)

func GetUserGroupIdUsingUnix(fileInfoWithPath *fileinfopath.Instance) *UserGroupId {
	if osconsts.IsWindows {
		return nil
	}

	return GetUserGroupIdUsing(fileInfoWithPath)
}
