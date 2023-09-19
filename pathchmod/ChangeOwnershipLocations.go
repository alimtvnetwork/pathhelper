package pathchmod

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
)

// ChangeOwnershipLocations not implemented
//
// TODO :
// https://gitlab.com/auk-go/pathhelper/-/issues/56
func ChangeOwnershipLocations(
	isContinueOnError,
	isRecursive bool,
	user, group string,
	errorCollection *errwrappers.Collection,
	locations []string,
) *errorwrapper.Wrapper {
	panic(errnew.NotImpl("https://gitlab.com/auk-go/pathhelper/-/issues/56"))
}
