package pathmodifier

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(
	isContinueOnErr bool,
	modifier *pathinsfmt.PathModifiersApply,
) *errorwrapper.Wrapper {
	if modifier == nil ||
		modifier.IsEmptyPathModifiers() ||
		modifier.IsEmptyGenericPathsCollection() {
		return errnew.EmptyPtr
	}

	errCollection := errwrappers.Empty()

	flatPaths := modifier.
		GenericPathsCollection.
		LazyFlatPaths()

	for _, pathModifier := range modifier.PathModifiers {
		errWrapper := ApplySimple(
			isContinueOnErr,
			&pathModifier,
			flatPaths)

		errCollection.AddWrapperPtr(errWrapper)
	}

	return errCollection.
		GetAsErrorWrapperPtr()
}
