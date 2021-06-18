package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ApplyChmod(
	isRecursive bool,
	isSkipOnNonExist bool,
	changeFileMode os.FileMode,
	location string,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	changingChmodRwxWrapper := chmodhelper.NewUsingFileMode(changeFileMode)

	if isRecursive {
		err := changingChmodRwxWrapper.LinuxApplyRecursive(isSkipOnNonExist, location)

		return &changingChmodRwxWrapper, errorwrapper.NewFilePtr(
			errtype.ChmodApplyFailed,
			err,
			location)
	}

	err := changingChmodRwxWrapper.ApplyChmod(isSkipOnNonExist, location)

	return &changingChmodRwxWrapper, errorwrapper.NewFilePtr(
		errtype.ChmodApplyFailed,
		err,
		location)
}
