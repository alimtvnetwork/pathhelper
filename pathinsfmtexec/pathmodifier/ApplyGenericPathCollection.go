package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyGenericPathCollection(
	isContinueOnErr bool,
	modifier *pathinsfmt.PathModifier,
	genericPathsCollection *pathinsfmt.GenericPathsCollection,
) *errorwrapper.Wrapper {
	if modifier == nil ||
		genericPathsCollection == nil ||
		genericPathsCollection.IsEmpty() {
		return errnew.EmptyPtr
	}

	return ApplySimple(
		isContinueOnErr,
		modifier,
		genericPathsCollection.LazyFlatPaths())
}
