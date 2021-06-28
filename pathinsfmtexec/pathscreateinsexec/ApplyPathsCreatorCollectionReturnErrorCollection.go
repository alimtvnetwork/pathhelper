package pathscreateinsexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyPathsCreatorCollectionReturnErrorCollection(
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
		isDeleteAllBeforeCreate,
		isLazyPaths,
		isIgnoreOnExist,
		errorCollection, pathsCreatorCollection)

	return errorCollection
}
