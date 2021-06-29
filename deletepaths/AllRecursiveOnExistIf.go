package deletepaths

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func AllRecursiveOnExistIf(
	isRemoveOnExistOnly bool,
	locations []string,
) *errorwrapper.Wrapper {
	if len(locations) == 0 {
		return errnew.EmptyPtr
	}

	if isRemoveOnExistOnly {
		return AllRecursiveOnExist(locations)
	}

	return AllRecursive(locations)
}
