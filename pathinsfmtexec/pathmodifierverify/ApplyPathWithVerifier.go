package pathmodifierverify

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyPathWithVerifier(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	pathWithVerifier *pathinsfmt.PathWithVerifier,
) (isSuccess bool) {
	if pathWithVerifier.IsVerifierUndefined() {
		return true
	}

	return ApplyVerifier(
		pathWithVerifier.IsNormalize,
		pathWithVerifier.IsRecursive,
		pathWithVerifier.IsSkipInvalid,
		isContinueOnError,
		pathWithVerifier.Verifier,
		errCollection,
		pathWithVerifier.CompiledPathAsSlice())
}
