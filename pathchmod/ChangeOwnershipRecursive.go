package pathchmod

import (
	"os/exec"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
)

func ChangeOwnershipRecursive(location, user, group string) *errorwrapper.Wrapper {
	_, err := exec.LookPath(constants.ChmodCommand)
	if err != nil || osconsts.IsWindows {
		return changeOwnershipWindowsRecursive(
			location,
			user,
			group)
	}

	return changeOwnershipUnixChmod(
		true,
		location,
		user,
		group)
}
