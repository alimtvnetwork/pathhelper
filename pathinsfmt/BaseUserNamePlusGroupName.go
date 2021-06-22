package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

type BaseUserNamePlusGroupName struct {
	BaseGroupName
	UserName *string `json:"UserName,omitempty"` // Not define or empty string or * means keeping the existing one
}

func (b *BaseUserNamePlusGroupName) IsUserNameExist() bool {
	return b != nil && !stringutil.IsNullOrEmptyPtr(b.UserName)
}

func (b *BaseUserNamePlusGroupName) UserNameSimple() string {
	return *b.UserName
}

func (b *BaseUserNamePlusGroupName) IsUsername(checkingUserName string) bool {
	isUsernameExist := b.IsUserNameExist()

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

func (b *BaseUserNamePlusGroupName) IsGroupNameExist() bool {
	return b != nil && b.GroupName != constants.EmptyString
}
