package pathmodifierverify

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyUsingFlatPathsDirectReturn(
	isContinueOnError bool,
	verifiers *pathinsfmt.PathVerifiers,
	locations []string,
) *errorwrapper.Wrapper {
	errCollection := errwrappers.Empty()

	isSuccess := ApplyUsingFlatPaths(
		isContinueOnError,
		verifiers,
		errCollection,
		locations)

	if !isSuccess {
		return errCollection.GetAsErrorWrapperPtr()
	}

	return nil
}
