package pathmodifier

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyGenericPathCollection(
	isContinueOnErr bool,
	errorCollection *errwrappers.Collection,
	modifier *pathinsfmt.PathModifier,
	genericPathsCollection *pathinsfmt.GenericPathsCollection,
) (isSuccess bool) {
	if modifier == nil ||
		genericPathsCollection == nil ||
		genericPathsCollection.IsEmpty() {
		return true
	}

	return ApplySimple(
		isContinueOnErr,
		errorCollection,
		modifier,
		genericPathsCollection.LazyFlatPaths())
}
