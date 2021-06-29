package deletepaths

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func All(locations []string) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return errnew.EmptyPtr
	}

	for _, location := range locations {
		recursiveErr := Single(location)

		if recursiveErr.HasError() {
			return recursiveErr
		}
	}

	return errnew.EmptyPtr
}

func AllOnExist(locations []string) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return errnew.EmptyPtr
	}

	for _, location := range locations {
		recursiveErr := SingleOnExist(location)

		if recursiveErr.HasError() {
			return recursiveErr
		}
	}

	return errnew.EmptyPtr
}
