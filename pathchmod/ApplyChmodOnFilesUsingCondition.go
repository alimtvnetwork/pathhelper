package pathchmod

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/errorwrapper"
)

func ApplyChmodOnFilesUsingCondition(
	condition *chmodins.Condition,
	changeFileMode os.FileMode,
	locations ...string,
) (*chmodins.RwxInstruction, *errorwrapper.Wrapper) {
	if len(locations) == 0 || condition == nil {
		return &chmodins.RwxInstruction{}, nil
	}

	return ApplyChmodOnFiles(
		condition.IsRecursive,
		condition.IsSkipOnInvalid,
		condition.IsContinueOnError,
		changeFileMode,
		locations...)
}
