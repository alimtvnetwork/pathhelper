package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/pathhelper/internal/cmdprefix"
)

func LinuxTouchFile(fullPath string) *errorwrapper.Wrapper {
	return errcmd.
		New.BashScript.ArgsErr(
		cmdprefix.Touch,
		fullPath,
	)
}
