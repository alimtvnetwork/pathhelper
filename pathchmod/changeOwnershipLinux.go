package pathchmod

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
)

func changeOwnershipLinux(path, user, group string) *errorwrapper.Wrapper {
	chownUserGroupArg := user + constants.Colon + group

	return errcmd.New.ShellScript.ArgsErr(
		constants.ChmodCommand,
		chownUserGroupArg,
		path,
	)
}
