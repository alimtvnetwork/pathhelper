package pathchmod

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ChmodChangeExecuteRevert(
	isRecursive,
	isSkipOnInvalid bool,
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
		isSkipOnInvalid,
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
		isSkipOnInvalid,
		location)

	return errnew.
		Path.
		Error(
			errtype.ChmodApplyFailed,
			err,
			location)
}
