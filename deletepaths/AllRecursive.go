package deletepaths

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func AllRecursive(locations []string) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return errnew.EmptyPtr
	}

	for _, location := range locations {
		recursiveErr := Recursive(location)

		if recursiveErr.HasError() {
			return recursiveErr
		}
	}

	return errnew.EmptyPtr
}
