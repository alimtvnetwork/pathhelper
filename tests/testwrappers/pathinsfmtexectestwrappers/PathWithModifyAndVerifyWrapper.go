package pathinsfmtexectestwrappers

import (
	"gitlab.com/auk-go/errorwrapper/errverify"

	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

type PathWithModifyAndVerifyWrapper struct {
	Header      string
	Modifier    *pathinsfmt.PathWithModifier
	Verifier    *pathinsfmt.PathWithVerifier
	ErrorVerify *errverify.CollectionVerifier
}
