package pathscreateinsexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/namegroup"
)

func applyUserNameGroupNameOnPathCreators(
	pathsCreator *pathinsfmt.PathsCreator,
	errorCollection *errwrappers.Collection,
) *errorwrapper.Wrapper {
	if pathsCreator == nil || pathsCreator.ApplyUserGroup == nil {
		return nil
	}

	errWrap := namegroup.Apply(
		true,
		false,
		pathsCreator.ApplyUserGroup,
		pathsCreator.RootDir)

	errorCollection.AddWrapperPtr(errWrap)

	return errWrap
}
