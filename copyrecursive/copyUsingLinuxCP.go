package copyrecursive

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func copyUsingLinuxCP(opts Options, src, dst string) *errorwrapper.Wrapper {
	if opts.IsRecursive {
		// run cp -r src dst
	} else {
		// run cp src dst
	}

	return errnew.NotImplemented
}
