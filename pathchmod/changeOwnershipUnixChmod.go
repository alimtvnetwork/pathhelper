package pathchmod

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
)

func changeOwnershipUnixChmod(
	isRecursive bool,
	location, user, group string,
) *errorwrapper.Wrapper {
	chmodCommandLine := cmdprefix.ChownUser(
		isRecursive,
		user,
		group)

	return errcmd.
		BashArgsErrorWrapper(
			chmodCommandLine,
			location)
}
