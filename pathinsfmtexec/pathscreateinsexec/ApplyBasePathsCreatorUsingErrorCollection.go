package pathscreateinsexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyBasePathsCreatorUsingErrorCollection(
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	errorCollection *errwrappers.Collection,
	basePathsCreator *pathinsfmt.BasePathsCreator,
) (
	isSuccess bool,
) {
	if basePathsCreator == nil {
		return true
	}

	pathsCreator := &pathinsfmt.PathsCreator{
		BasePathsCreator: *basePathsCreator,
		ApplyRwx:         nil,
		ApplyUserGroup:   nil,
	}

	return ApplyPathsCreatorUsingErrorCollection(
		isDeleteAllBeforeCreate,
		isLazyPaths,
		isIgnoreOnExist,
		errorCollection,
		pathsCreator)
}
