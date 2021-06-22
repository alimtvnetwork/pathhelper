package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
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

func (b *BaseUserNamePlusGroupName) HasUserName() bool {
	return b != nil && !stringutil.IsNullOrEmptyPtr(b.UserName)
}

func (b *BaseUserNamePlusGroupName) IsUserNameEmpty() bool {
	return b != nil || stringutil.IsNullOrEmptyPtr(b.UserName)
}

func (b *BaseUserNamePlusGroupName) IsGroupNameEmpty() bool {
	return b != nil || stringutil.IsEmptyOrWhitespace(b.GroupName)
}

func (b *BaseUserNamePlusGroupName) UserNameSimple() string {
	return *b.UserName
}

func (b *BaseUserNamePlusGroupName) IsUsername(checkingUserName string) bool {
	isUsernameExist := b.HasUserName()

	if !isUsernameExist && checkingUserName == "" {
		return true
	}

	if !isUsernameExist {
		return false
	}

	return checkingUserName == *b.UserName
}

func (b *BaseUserNamePlusGroupName) IsGroupName(checkingGroupName string) bool {
	return checkingGroupName == b.GroupName
}

func (b *BaseUserNamePlusGroupName) HasGroupName() bool {
	return b != nil && b.GroupName != constants.EmptyString
}

func (b *BaseUserNamePlusGroupName) IsGroupNameUserNameBothEmpty() bool {
	return b == nil || b.IsUserNameEmpty() && b.IsGroupNameEmpty()
}
