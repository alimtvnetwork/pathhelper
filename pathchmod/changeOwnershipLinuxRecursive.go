package pathchmod

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/argsinternal"
)

func changeOwnershipUnixChmodRecursive(path, user, group string) *errorwrapper.Wrapper {
	chownUserGroupArg := user + constants.Colon + group
	chmodCommandLine := argsinternal.Join(
		constants.ChmodCommand,
		constants.RecursiveCommandFlag,
		chownUserGroupArg,
		path)

	return errcmd.
		ShellScriptsErrorWrapper(chmodCommandLine)
}
