package pathchmod

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/pathhelper/internal/cmdprefix"
)

func changeOwnershipUnixChmod(
	isRecursive bool,
	location, user, group string,
) *errorwrapper.Wrapper {
	chownUser := cmdprefix.ChownUser(
		isRecursive,
		user,
		group)

	return errcmd.
		New.
		BashScript.
		ArgsErr(
			chownUser,
			location)
}
