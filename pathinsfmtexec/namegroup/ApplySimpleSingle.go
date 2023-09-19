package namegroup

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathchmod"
)

func ApplySimpleSingle(
	isRecursive bool,
	userName string,
	groupName string,
	location string,
) *errorwrapper.Wrapper {
	return pathchmod.ChangeOwnershipOptions(
		isRecursive,
		location,
		userName,
		groupName)
}
