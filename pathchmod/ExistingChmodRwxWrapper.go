package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ExistingChmodRwxWrapper(
	location string,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	existingChmod, err := chmodhelper.GetExistingChmodRwxWrapperPtr(location)

	if err != nil {
		return existingChmod, errorwrapper.NewPath(
			errtype.ExistingChmodReadFailed,
			err,
			location)
	}

	return existingChmod, errnew.EmptyPtr
}
