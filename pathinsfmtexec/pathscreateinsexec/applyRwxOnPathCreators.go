package pathscreateinsexec

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathchmod"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func applyRwxOnPathCreators(
	pathsCreator *pathinsfmt.PathsCreator,
	errorCollection *errwrappers.Collection,
) *errorwrapper.Wrapper {
	if pathsCreator == nil || pathsCreator.ApplyRwx == nil {
		return nil
	}

	if osconsts.IsWindows {
		return nil
	}

	fileMode, errWrap := pathchmod.ParseRwxOwnerGroupOtherToFileMode(
		pathsCreator.ApplyRwx)

	errorCollection.AddWrapperPtr(errWrap)

	if errWrap.HasError() {
		return errWrap
	}

	chmodApplyErr := pathchmod.ApplyLinuxRecursiveChmodOnPathUsingFileMode(
		fileMode,
		pathsCreator.RootDir)
	errorCollection.AddWrapperPtr(chmodApplyErr)

	return chmodApplyErr
}
