package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ApplyChmodRwxInstructions(
	instructions *chmodins.BaseRwxInstructions,
	paths []string,
) *errorwrapper.Wrapper {
	if instructions == nil ||
		instructions.RwxInstructions == nil ||
		len(paths) == 0 {
		return nil
	}

	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		instructions.RwxInstructions)

	if err != nil {
		return errnew.Type.Error(
			errtype.ParsingFailed,
			err)
	}

	err2 := executors.ApplyOnPathsPtr(
		&paths)

	if err2 != nil {
		return errnew.Type.Error(
			errtype.ChmodApplyFailed,
			err2)
	}

	return nil
}
