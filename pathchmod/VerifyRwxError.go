package pathchmod

import (
	"strings"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func VerifyRwxErrorLocations(
	isRecursiveErrorIgnore bool,
	instruction *chmodins.RwxInstruction,
	locations []string,
) *errorwrapper.Wrapper {
	if len(locations) == 0 || instruction == nil {
		return nil
	}

	executor, err := chmodhelper.
		ParseRwxInstructionToExecutor(instruction)

	if err != nil {
		return errnew.
			Path.
			Error(
				errtype.ChmodInvalid,
				err,
				strings.Join(locations, constants.CommaSpace))
	}

	err2 := executor.VerifyRwxModifiers(isRecursiveErrorIgnore, locations)

	return errnew.Type.Error(errtype.ChmodMismatch, err2)
}
