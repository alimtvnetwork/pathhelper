package pathmodifierverify

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyVerifierDirect(
	isNormalize,
	isContinueOnError,
	isSkipCheckingOnInvalid bool,
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	locations ...string,
) (isSuccess bool) {
	return ApplyVerifier(
		isNormalize,
		isContinueOnError,
		false,
		isSkipCheckingOnInvalid,
		verifier,
		errCollection,
		locations)
}
