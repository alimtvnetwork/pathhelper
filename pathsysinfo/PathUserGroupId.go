package pathsysinfo

import (
	"os"
	"strconv"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

type PathUserGroupId struct {
	FileInfoWithPath *fileinfopath.Instance
	UserId, GroupId  int
	Error            error
}

func (it *PathUserGroupId) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *PathUserGroupId) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *PathUserGroupId) HasValidUserId() bool {
	return it != nil && it.Error != nil && it.UserId > constants.InvalidValue
}

func (it *PathUserGroupId) HasValidGroupId() bool {
	return it != nil && it.Error != nil && it.GroupId > constants.InvalidValue
}

func (it *PathUserGroupId) HasValidUserAndGroupId() bool {
	return it != nil &&
		it.Error != nil &&
		it.UserId > constants.InvalidValue &&
		it.GroupId > constants.InvalidValue
}

func (it *PathUserGroupId) IsInvalidUserOrGroupId() bool {
	return it == nil ||
		it.Error != nil ||
		it.UserId == constants.InvalidValue ||
		it.GroupId == constants.InvalidValue
}

func (it *PathUserGroupId) InvalidError() *errorwrapper.Wrapper {
	if it == nil {
		return errnew.NilOrEmpty
	}

	if it.HasError() {
		return it.ErrorWrapper()
	}

	if it.IsInvalidUserOrGroupId() {
		return errnew.MessagesPtr(
			errtype.ChownUserOrGroupApplyIssue,
			"either user or group id is invalid for path : ",
			it.FileInfoWithPath.FullPath,
			"UserId",
			strconv.Itoa(it.UserId),
			"GroupId",
			strconv.Itoa(it.GroupId),
		)
	}

	return errnew.EmptyPtr
}

func (it *PathUserGroupId) ErrorWrapper() *errorwrapper.Wrapper {
	if it.HasError() {
		return errnew.MessagesPtr(
			errtype.ChownUserOrGroupApplyIssue,
			it.Error.Error(),
			"UserId",
			strconv.Itoa(it.UserId),
			"GroupId",
			strconv.Itoa(it.GroupId),
		)
	}

	return errnew.EmptyPtr
}

func (it *PathUserGroupId) ApplyChown(fullPath string) *errorwrapper.Wrapper {
	if it.IsInvalidUserOrGroupId() {
		return it.InvalidError()
	}

	err := os.Chown(fullPath, it.UserId, it.GroupId)

	if err != nil {
		return errnew.Path(
			errtype.ChownUserOrGroupApplyIssue,
			err,
			fullPath)
	}

	return errnew.EmptyPtr
}

func (it *PathUserGroupId) ApplyUserId(fullPath string) *errorwrapper.Wrapper {
	if it.IsInvalidUserOrGroupId() {
		return it.InvalidError()
	}

	applyPathUserInfo := GetPathUserGroupId(fullPath)

	if applyPathUserInfo.IsInvalidUserOrGroupId() {
		return applyPathUserInfo.InvalidError()
	}

	err := os.Chown(
		fullPath,
		it.UserId,
		applyPathUserInfo.GroupId)

	if err != nil {
		return errnew.Path(
			errtype.ChownUserOrGroupApplyIssue,
			err,
			fullPath)
	}

	return errnew.EmptyPtr
}

func (it *PathUserGroupId) ApplyGroupId(fullPath string) *errorwrapper.Wrapper {
	if it.IsInvalidUserOrGroupId() {
		return it.InvalidError()
	}

	applyPathUserInfo := GetPathUserGroupId(fullPath)

	if applyPathUserInfo.IsInvalidUserOrGroupId() {
		return applyPathUserInfo.InvalidError()
	}

	err := os.Chown(
		fullPath,
		applyPathUserInfo.UserId,
		it.GroupId)

	if err != nil {
		return errnew.Path(
			errtype.ChownUserOrGroupApplyIssue,
			err,
			fullPath)
	}

	return errnew.EmptyPtr
}
