package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func ApplyChmodOnFiles(
	isRecursive,
	isSkipOnInvalid,
	isContinueOnError bool,
	changeFileMode os.FileMode,
	locations ...string,
) (*chmodins.RwxInstruction, *errorwrapper.Wrapper) {
	if len(locations) == 0 {
		return &chmodins.RwxInstruction{}, errnew.EmptyPtr
	}

	changingChmodRwxWrapper := chmodhelper.NewUsingFileMode(changeFileMode)
	rwxOwnerGroupOther := changingChmodRwxWrapper.ToRwxOwnerGroupOther()
	rwxInstruction := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition: chmodins.Condition{
			IsSkipOnInvalid:   isSkipOnInvalid,
			IsContinueOnError: isContinueOnError,
			IsRecursive:       isRecursive,
		},
	}

	executor, err := chmodhelper.ParseRwxInstructionToExecutor(rwxInstruction)

	if err != nil {
		return rwxInstruction, errnew.NewPtr(errtype.Conversion, err)
	}

	err2 := executor.ApplyOnPaths(locations)

	return rwxInstruction, errorwrapper.NewRefs(
		errtype.ChmodApplyFailed,
		err2,
		ref.Value{
			Variable: "locations",
			Value:    locations,
		})
}
