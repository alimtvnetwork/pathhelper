package pathchmod

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func changeOwnershipUnixChmodRecursive(path, user, group string) *errorwrapper.Wrapper {
	chownUserGroupArg := user + constants.Colon + group
	chmodCommandLine := errcmd.ArgsJoin(
		constants.ChmodCommand,
		constants.RecursiveCommandFlag,
		chownUserGroupArg,
		path)

	return errcmd.
		ShellScriptsErrorWrapper(chmodCommandLine)
}
