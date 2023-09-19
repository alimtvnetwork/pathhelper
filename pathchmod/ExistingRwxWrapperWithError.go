package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ExistingRwxWrapperWithError(location string) *RwxWrapperWithError {
	if location == "" {
		return &RwxWrapperWithError{
			RwxWrapper: nil,
			ErrWrapper: errnew.Path.Empty(),
		}
	}

	rwxWrapper, err := chmodhelper.GetExistingChmodRwxWrapperPtr(location)
	if err != nil {
		pathErr := errnew.
			Path.
			Messages(
				errtype.File,
				location,
				"ExistingRwxWrapperWithError",
				err.Error())

		return &RwxWrapperWithError{
			RwxWrapper: nil,
			ErrWrapper: pathErr,
		}
	}

	return &RwxWrapperWithError{
		RwxWrapper: rwxWrapper,
		ErrWrapper: nil,
	}
}
