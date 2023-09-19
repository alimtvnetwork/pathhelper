package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ExistingChmodWithError(location string) *ChmodWithError {
	if location == "" {
		return &ChmodWithError{
			ErrWrapper: errnew.Path.Empty(),
		}
	}

	chmod, err := chmodhelper.GetExistingChmod(location)
	if err != nil {
		pathErr := errnew.
			Path.
			Messages(
				errtype.File,
				location,
				"ExistingChmodWithError",
				err.Error())

		return &ChmodWithError{
			Chmod:      0,
			ErrWrapper: pathErr,
		}
	}

	return &ChmodWithError{
		Chmod:      chmod,
		ErrWrapper: nil,
	}
}
