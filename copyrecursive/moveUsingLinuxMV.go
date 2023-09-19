package copyrecursive

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
)

func moveUsingLinuxMV(opts Options, src, dst string) *errorwrapper.Wrapper {
	// Options mapping for mv command
	// 	IsSkipOnExist:  -n
	// 	IsClearDestination:     --remove-destination,
	args := make([]string, 0, constants.ArbitraryCapacity8)
	if opts.IsSkipOnExist {
		args = append(args, mvArgSkipOnExist)
	}

	args = append(args, src)
	args = append(args, dst)

	return errcmd.New.Create(
		false,
		false,
		mvCommand, args...,
	).CompiledErrorWrapper()
}
