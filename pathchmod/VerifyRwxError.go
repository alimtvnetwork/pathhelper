package pathchmod

import (
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func VerifyRwxErrorLocations(
	isRecursiveErrorIgnore bool,
	instruction *chmodins.RwxInstruction,
	locations []string,
) *errorwrapper.Wrapper {
	if len(locations) == 0 || instruction == nil {
		return errnew.EmptyPtr
	}

	executor, err := chmodhelper.
		ParseRwxInstructionToExecutor(instruction)

	if err != nil {
		return errnew.Path(
			errtype.ChmodInvalid,
			err,
			strings.Join(locations, constants.CommaSpace))
	}

	err2 := executor.VerifyRwxModifiers(isRecursiveErrorIgnore, locations)

	return errnew.NewPtr(errtype.ChmodMismatch, err2)
}
