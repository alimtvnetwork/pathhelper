package pathscreateinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func applyRwxOnPathCreators(
	pathsCreator *pathinsfmt.PathsCreator,
	errorCollection *errwrappers.Collection,
) *errorwrapper.Wrapper {
	if pathsCreator == nil || pathsCreator.ApplyRwx == nil {
		return nil
	}

	fileMode, errWp := pathchmod.ParseRwxOwnerGroupOtherToFileMode(
		pathsCreator.ApplyRwx)

	errorCollection.AddWrapperPtr(errWp)

	if errWp.HasError() {
		return errWp
	}

	chmodApplyErr := pathchmod.ApplyLinuxRecursiveChmodOnPathUsingFileMode(
		fileMode,
		pathsCreator.RootDir)
	errorCollection.AddWrapperPtr(chmodApplyErr)

	return chmodApplyErr
}
