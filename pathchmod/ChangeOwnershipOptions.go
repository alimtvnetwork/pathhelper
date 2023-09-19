package pathchmod

import "gitlab.com/auk-go/errorwrapper"

func ChangeOwnershipOptions(
	isRecursive bool,
	location, user, group string,
) *errorwrapper.Wrapper {
	if isRecursive {
		return ChangeOwnershipRecursive(location, user, group)
	}

	return ChangeOwnership(location, user, group)
}
