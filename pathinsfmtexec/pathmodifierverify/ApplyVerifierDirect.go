package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyVerifierDirect(
	isNormalize,
	isContinueOnError,
	isSkipCheckingOnInvalid bool,
	verifier *pathinsfmt.PathVerifier,
	locations ...string,
) *errorwrapper.Wrapper {
	return ApplyVerifier(
		isNormalize,
		isContinueOnError,
		false,
		isSkipCheckingOnInvalid,
		verifier,
		locations)
}
