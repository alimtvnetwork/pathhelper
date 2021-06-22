package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyVerifier(
	verifier *pathinsfmt.PathVerifier,
	locations []string,
) *errorwrapper.Wrapper {
	if verifier == nil || len(locations) == 0 {
		return errnew.EmptyPtr
	}

	errCollection := errwrappers.Empty()

	for _, location := range locations {
		workingPath := normalize.PathUsingSingleIf(
			verifier.IsNormalize,
			location)

		if verifier.IsRecursiveCheck {
			collectRecursiveCheckErrors(errCollection, verifier, workingPath)
		} else {
			// non recursive
			errWp := applyVerifierSinglePathNonRecursive(
				true,
				false,
				verifier,
				workingPath)

			errCollection.AddWrapperPtr(errWp)
		}
	}

	return errCollection.GetAsErrorWrapperPtr()
}
