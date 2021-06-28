package pathscreateinsexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyPathsCreatorCollectionsReturnErrorCollection(
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	pathsCreatorCollection []*pathinsfmt.PathsCreatorCollection,
) (
	errorCollection *errwrappers.Collection,
) {
	if pathsCreatorCollection == nil {
		return errwrappers.Empty()
	}

	errorCollection = errwrappers.Empty()

	for _, collection := range pathsCreatorCollection {
		ApplyPathsCreatorCollectionUsingErrorCollection(
			isDeleteAllBeforeCreate,
			isLazyPaths,
			isIgnoreOnExist,
			errorCollection,
			collection)
	}

	return errorCollection
}
