package deletepaths

import (
	"gitlab.com/auk-go/errorwrapper"
)

func AllRecursiveOnExistIf(
	isRemoveOnExistOnly bool,
	locations []string,
) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return nil
	}

	if isRemoveOnExistOnly {
		return AllRecursiveOnExist(locations)
	}

	return AllRecursive(locations)
}
