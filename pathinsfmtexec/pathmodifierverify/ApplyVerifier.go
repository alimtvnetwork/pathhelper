package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
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

		isFileExist := fsinternal.IsPathExists(workingPath)
		isFileMissing := !isFileExist

		if verifier.IsSkipCheckingOnNonExist && isFileMissing {
			return errnew.EmptyPtr
		}

		if !verifier.IsSkipCheckingOnNonExist && isFileMissing {
			return errnew.PathMessages(
				errtype.PathNotFound,
				workingPath,
				"Use IsSkipCheckingOnNonExist to true skip the error.")
		}

		if verifier.IsRecursiveCheck {
			collectRecursiveCheckErrors(errCollection, verifier, location)
		} else {
			// non recursive
			errWp := applyVerifierSinglePathNonRecursive(verifier, location)

			errCollection.AddWrapperPtr(errWp)
		}
	}

	return errCollection.GetAsErrorWrapperPtr()
}
