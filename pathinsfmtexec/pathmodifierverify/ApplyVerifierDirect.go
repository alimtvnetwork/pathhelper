package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyVerifierDirect(
	verifier *pathinsfmt.PathVerifier,
	locations ...string,
) *errorwrapper.Wrapper {
	return ApplyVerifier(verifier, locations)
}
