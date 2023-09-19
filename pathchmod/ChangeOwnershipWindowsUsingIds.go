package pathchmod

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

// ChangeOwnershipWindowsUsingIds non recursive
func ChangeOwnershipWindowsUsingIds(
	location string,
	userId, groupId int,
) *errorwrapper.Wrapper {
	err := os.Chown(location, userId, groupId)

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
