package copyrecursive

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func UsingInstructions(
	instruction *Instruction,
) *errorwrapper.Wrapper {
	if instruction == nil {
		// more verbose output
		return errnew.InvalidInput
	}

	return nil
}
