package pathchmod

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

// ChangeOwnershipWindows non-recursive
func ChangeOwnershipWindows(location, user, group string) *errorwrapper.Wrapper {
	uid, gid, errorWrapper := GetUserGroupId(user, group)
	if errorWrapper.HasError() {
		return errorWrapper
	}

	err := os.Chown(location, uid, gid)

	if err != nil {
		return errnew.
			Path.
			Error(
				errtype.ChownUserOrGroupApplyIssue,
				err,
				location)
	}

	return nil
}
