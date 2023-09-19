package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/pathhelper/internal/cmdprefix"
)

func LinuxTouchFileUsingLock(fullPath string) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return errcmd.
		New.BashScript.ArgsErr(
		cmdprefix.Touch,
		fullPath,
	)
}
