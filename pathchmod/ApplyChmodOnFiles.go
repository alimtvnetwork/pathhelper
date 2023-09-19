package pathchmod

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/ref"
)

func ApplyChmodOnFiles(
	isRecursive,
	isSkipOnInvalid,
	isContinueOnError bool,
	changeFileMode os.FileMode,
	locations ...string,
) (*chmodins.RwxInstruction, *errorwrapper.Wrapper) {
	if len(locations) == 0 {
		return &chmodins.RwxInstruction{}, nil
	}

	changingChmodRwxWrapper := chmodhelper.New.RwxWrapper.UsingFileMode(changeFileMode)
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
		return rwxInstruction, errnew.Type.Error(errtype.Conversion, err)
	}

	finalErr := executor.ApplyOnPaths(locations)

	if finalErr == nil {
		return rwxInstruction, nil
	}

	return rwxInstruction, errorwrapper.NewRefs(
		codestack.SkipNone,
		errtype.ChmodApplyFailed,
		finalErr,
		ref.Value{
			Variable: "locations",
			Value:    locations,
		})
}
