package pathsysinfo

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/fileinfopath"
)

func GetPathUserGroupIdUsing(instance *fileinfopath.Instance) *PathUserGroupId {
	return &PathUserGroupId{
		FileInfoWithPath: instance,
		UserId:           constants.InvalidValue,
		GroupId:          constants.InvalidValue,
		Error: errtype.NotSupportInWindows.Error(
			constants.EmptyString,
			"path",
			instance.FullPath),
	}
}
