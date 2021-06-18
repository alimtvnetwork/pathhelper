package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ChmodChangeExecuteRevert(
	isRecursive,
	isSkipOnNonExist bool,
	changeFileMode os.FileMode,
	location string,
	executor func(location string) *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	existingChmod, errWrapper := ExistingChmodRwxWrapper(location)

	if errWrapper.HasError() {
		return errWrapper
	}

	_, rwxErrorWrapper := ApplyChmod(
		isRecursive,
		isSkipOnNonExist,
		changeFileMode,
		location)

	if rwxErrorWrapper.HasError() {
		return rwxErrorWrapper
	}

	executionErr := executor(location)

	if executionErr.HasError() {
		return executionErr
	}

	err := existingChmod.ApplyChmod(
		isSkipOnNonExist,
		location)

	return errorwrapper.NewFilePtr(
		errtype.ChmodApplyFailed,
		err,
		location)
}
