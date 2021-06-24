package pathsysinfo

import "gitlab.com/evatix-go/errorwrapper/errnew"

func GetGroupInfo(groupName string) *GroupInfo {
	groupObj, errWrapper := LookupGroup(groupName)

	if errWrapper.HasError() {
		return InvalidGroupInfo(errWrapper)
	}

	groupId, groupIdConvertFailed := GetGroupId(groupObj)

	if groupIdConvertFailed.HasError() {
		return &GroupInfo{
			Group:        groupObj,
			Id:           groupId,
			IsValidGroup: true,
			HasValidId:   false,
			ErrorWrapper: groupIdConvertFailed,
		}
	}

	return &GroupInfo{
		Group:        groupObj,
		Id:           groupId,
		IsValidGroup: true,
		HasValidId:   true,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
