package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyLocationCollection(
	isContinueOnErr bool,
	modifier *pathinsfmt.PathModifier,
	locationCollection *pathinsfmt.LocationCollection,
) *errorwrapper.Wrapper {
	if modifier == nil ||
		locationCollection == nil ||
		locationCollection.IsEmpty() {
		return errnew.EmptyPtr
	}

	return ApplySimple(
		isContinueOnErr,
		modifier,
		locationCollection.LazyFlatPaths())
}
