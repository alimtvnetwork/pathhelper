package pathchmod

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ApplyChmodDefault(
	changeFileMode os.FileMode,
	location string,
) *errorwrapper.Wrapper {
	err := chmodhelper.ChmodApply.Default(
		changeFileMode,
		location)

	if err == nil {
		return nil
	}

	return errnew.Error.Default(
		errtype.ChmodApplyFailed,
		err)
}
