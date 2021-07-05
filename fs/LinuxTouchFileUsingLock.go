package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
)

func LinuxTouchFileUsingLock(fullPath string) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return errcmd.BashArgsErrorWrapper(
		cmdprefix.Touch,
		fullPath,
	)
}
