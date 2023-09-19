package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
)

func ChmodChangeExecuteRevertUsingRwxWrapper(
	isRecursive,
	isSkipOnInvalid bool,
	changeFileModeRwxWrapper *chmodhelper.RwxWrapper,
	location string,
	executor func(location string) *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	return ChmodChangeExecuteRevert(
		isRecursive,
		isSkipOnInvalid,
		changeFileModeRwxWrapper.ToFileMode(),
		location,
		executor)
}
