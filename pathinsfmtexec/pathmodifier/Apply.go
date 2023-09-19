package pathmodifier

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func Apply(
	isContinueOnErr bool,
	modifier *pathinsfmt.PathModifiersApply,
) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	if modifier == nil ||
		modifier.IsEmptyPathModifiers() ||
		modifier.IsEmptyGenericPathsCollection() {
		return errCollection
	}

	ApplyUsingErrorCollection(
		isContinueOnErr,
		errCollection,
		modifier)

	return errCollection
}
