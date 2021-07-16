package copyrecursive

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type Instruction struct {
	SourceDestination
	Options
}

func (it *Instruction) Run() *errorwrapper.Wrapper {
	if it.IsUseShellOrCmd {
		return copyUsingLinuxCP(it.Options, it.Source, it.Destination)
	}

	return errnew.NotImplemented
}
