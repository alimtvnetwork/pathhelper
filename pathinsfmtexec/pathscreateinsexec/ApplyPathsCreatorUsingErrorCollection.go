package pathscreateinsexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
)

func ApplyPathsCreatorUsingErrorCollection(
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	errorCollection *errwrappers.Collection,
	pathsCreator *pathinsfmt.PathsCreator,
) (
	isSuccess bool,
) {
	if pathsCreator == nil {
		return true
	}

	var workingPaths []string

	if isLazyPaths {
		workingPaths = pathsCreator.LazyFlatPaths()
	} else {
		workingPaths = pathsCreator.FlatPaths()
	}

	errCount := errorCollection.Length()

	if isDeleteAllBeforeCreate {
		errorCollection.AddAnyFunctions(
			pathsCreator.DeleteAllPaths)
	}

	// paths create
	if pathsCreator.HasRwx() {
		fileMode, errWp := pathchmod.ParseRwxOwnerGroupOtherToFileMode(pathsCreator.ApplyRwx)

		errorCollection.AddWrapperPtr(errWp)

		if errWp.HasError() {
			return false
		}

		_, filesCreateErr := createpath.CreateManySameDirWithFileMode(
			fileMode,
			isIgnoreOnExist,
			pathsCreator.RootDir,
			workingPaths)

		errorCollection.AddWrapperPtr(filesCreateErr)

		if filesCreateErr.HasError() {
			return false
		}
	} else {
		// create without chmod
		_, filesCreateErr := createpath.CreateMany(
			isIgnoreOnExist,
			workingPaths)

		errorCollection.AddWrapperPtr(filesCreateErr)

		if filesCreateErr.HasError() {
			return false
		}
	}

	// apply groups
	if pathsCreator.HasUserGroup() {
		errWp := namegroup.Apply(
			true,
			false,
			pathsCreator.ApplyUserGroup,
			pathsCreator.RootDir)

		errorCollection.AddWrapperPtr(errWp)
	}

	return errCount == errorCollection.Length()
}
