package pathsysinfo

import (
	"syscall"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/fileinfopath"
)

func GetPathUserGroupIdUsing(instance *fileinfopath.Instance) *PathUserGroupId {
	if instance.IsInvalidPath() {
		return &PathUserGroupId{
			FileInfoWithPath: instance,
			UserId:           constants.InvalidValue,
			GroupId:          constants.InvalidValue,
			Error: errtype.InvalidPath.Error(
				instance.Error.Error(),
				"path",
				instance.FullPath),
		}
	}

	fileSys := instance.FileInfo.Sys()
	stat, isOkay := fileSys.(*syscall.Stat_t)

	if isOkay {
		uid := stat.Uid
		gid := stat.Gid

		return &PathUserGroupId{
			FileInfoWithPath: instance,
			UserId:           int(uid),
			GroupId:          int(gid),
			Error:            nil,
		}
	}

	return &PathUserGroupId{
		FileInfoWithPath: instance,
		UserId:           constants.InvalidValue,
		GroupId:          constants.InvalidValue,
		Error: errtype.StatFailed.Error(
			"couldn't retrieve userid, group id from path",
			"path",
			instance.FullPath),
	}
}
