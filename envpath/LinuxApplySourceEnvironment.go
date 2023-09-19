package envpath

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
)

func LinuxApplySourceEnvironment() *errorwrapper.Wrapper {
	return errcmd.New.BashScript.ArgsDefault("source", etcEnvPath).
		CompiledResult().
		ErrorWrapper()
}
