package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ExistingChmodRwxWrapper(
	location string,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	existingChmod, err := chmodhelper.GetExistingChmodRwxWrapperPtr(location)

	if err != nil {
		return existingChmod, errnew.
			Path.
			Error(
				errtype.ExistingChmodReadFailed,
				err,
				location)
	}

	return existingChmod, nil
}
