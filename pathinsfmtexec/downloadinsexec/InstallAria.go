package downloadinsexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
)

func InstallAria() *errorwrapper.Wrapper {
	return errcmd.
		New.BashScript.ArgsErr(
		installAriaBash)
}
