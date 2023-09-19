package pathmodifierverify

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyVerifierUsingOption(
	errCollection *errwrappers.Collection,
	options *VerifierOption,
	verifier *pathinsfmt.PathVerifier,
	locations []string,
) (isSuccess bool) {
	return ApplyVerifier(
		options.IsNormalize,
		options.IsRecursive,
		options.IsSkipInvalid,
		options.IsContinueOnError,
		verifier,
		errCollection,
		locations)
}
