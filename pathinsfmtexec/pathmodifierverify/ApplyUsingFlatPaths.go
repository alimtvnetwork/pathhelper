package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingFlatPaths(
	verifiers *pathinsfmt.PathVerifiers,
	locations []string,
) *errorwrapper.Wrapper {
	if verifiers == nil || verifiers.IsEmpty() {
		return errnew.EmptyPtr
	}

	for _, verifier := range verifiers.PathVerifiers {
		errWp := ApplyVerifier(&verifier, locations)

		if errWp.HasError() {
			return errWp
		}
	}

	return errnew.EmptyPtr
}
