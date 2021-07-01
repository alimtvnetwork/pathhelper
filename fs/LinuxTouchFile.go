package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
)

func LinuxTouchFile(fullPath string) *errorwrapper.Wrapper {
	return errcmd.BashArgsErrorWrapper(
		cmdprefix.Touch,
		fullPath,
	)
}
