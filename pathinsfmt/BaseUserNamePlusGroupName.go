package pathinsfmt

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

type BaseUserNamePlusGroupName struct {
	BaseGroupName
	UserName *string `json:"UserName,omitempty"` // Not define or empty string or * means keeping the existing one
}

func NewBaseUserNamePlusGroupName(
	username, groupName string,
) *BaseUserNamePlusGroupName {
	return &BaseUserNamePlusGroupName{
		BaseGroupName: BaseGroupName{
			GroupName: groupName,
		},
		UserName: &username,
	}
}

func (it *BaseUserNamePlusGroupName) HasUserNameOrGroup() bool {
	return it.HasUserName() || it.HasGroupName()
}

func (it *BaseUserNamePlusGroupName) HasUserName() bool {
	return it != nil && !stringutil.IsNullOrEmptyPtr(it.UserName)
}

func (it *BaseUserNamePlusGroupName) IsUserNameEmpty() bool {
	return it == nil || stringutil.IsNullOrEmptyPtr(it.UserName)
}

func (it *BaseUserNamePlusGroupName) UserNameSimple() string {
	return *it.UserName
}

func (it *BaseUserNamePlusGroupName) IsUsername(checkingUserName string) bool {
	isUsernameExist := it.HasUserName()

	if !isUsernameExist && checkingUserName == "" {
		return true
	}

	if !isUsernameExist {
		return false
	}

	return checkingUserName == *it.UserName
}

func (it *BaseUserNamePlusGroupName) IsGroupNameUserNameBothEmpty() bool {
	return it == nil || it.IsUserNameEmpty() && it.IsGroupNameEmpty()
}

func (it *BaseUserNamePlusGroupName) Clone() *BaseUserNamePlusGroupName {
	if it == nil {
		return nil
	}

	var userName *string

	if it.UserName != nil {
		userName2 := *it.UserName
		userName = &userName2
	}

	return &BaseUserNamePlusGroupName{
		BaseGroupName: BaseGroupName{
			GroupName: it.GroupName,
		},
		UserName: userName,
	}
}
