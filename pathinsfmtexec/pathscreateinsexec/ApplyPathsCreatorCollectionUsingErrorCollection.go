package pathscreateinsexec

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
)

func ApplyPathsCreatorCollectionUsingErrorCollection(
	isDeleteAllBeforeCreate,
	isLazyPaths,
	isIgnoreOnExist bool,
	errorCollection *errwrappers.Collection,
	pathsCreatorCollection *pathinsfmt.PathsCreatorCollection,
) (
	isSuccess bool,
) {
	if pathsCreatorCollection == nil {
		return true
	}

	errCount := errorCollection.Length()

	if isDeleteAllBeforeCreate {
		for _, instruction := range pathsCreatorCollection.PathsCreateInstructions {
			errorCollection.AddAnyFunctions(instruction.DeleteAllPaths)
		}
	}

	workingPaths := pathsCreatorCollection.
		LazyFlatPathsIf(isLazyPaths)

	// paths create
	_, filesCreateErr := createpath.CreateMany(
		isIgnoreOnExist,
		workingPaths)

	errorCollection.AddWrapperPtr(filesCreateErr)

	if filesCreateErr.HasError() {
		return false
	}

	// apply chmod
	if pathsCreatorCollection.HasRwx() {
		fileMode, errWp := pathchmod.ParseRwxOwnerGroupOtherToFileMode(
			pathsCreatorCollection.ApplyRwx)

		errorCollection.AddWrapperPtr(errWp)

		if errWp.HasError() {
			return false
		}

		for _, instruction := range pathsCreatorCollection.PathsCreateInstructions {
			instruction.ApplyLinuxRecursiveFileModeOnRoot(fileMode)
		}

		errorCollection.AddWrapperPtr(filesCreateErr)

		if filesCreateErr.HasError() {
			return false
		}
	}

	// apply groups
	if pathsCreatorCollection.HasUserGroup() && osconsts.IsUnixGroup {
		for _, instruction := range pathsCreatorCollection.PathsCreateInstructions {
			errWp := namegroup.Apply(
				true,
				false,
				pathsCreatorCollection.ApplyUserGroup,
				instruction.RootDir)

			errorCollection.AddWrapperPtr(errWp)
		}
	}

	return errCount == errorCollection.Length()
}
