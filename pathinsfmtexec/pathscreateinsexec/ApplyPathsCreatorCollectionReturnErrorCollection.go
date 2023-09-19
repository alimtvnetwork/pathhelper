package pathscreateinsexec

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyPathsCreatorCollectionReturnErrorCollection(
	isLock,
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	pathsCreatorCollection *pathinsfmt.PathsCreatorCollection,
) (
	errorCollection *errwrappers.Collection,
) {
	if pathsCreatorCollection == nil {
		return errwrappers.Empty()
	}

	errorCollection = errwrappers.Empty()

	ApplyPathsCreatorCollectionUsingErrorCollection(
		isLock,
		isDeleteAllBeforeCreate,
		isLazyPaths,
		isIgnoreOnExist,
		errorCollection, pathsCreatorCollection)

	return errorCollection
}
