package pathinsfmt

import "gitlab.com/evatix-go/core/coreinstruction"

type PathVerifiers struct {
	coreinstruction.BaseSpecPlusRequestIds
	PathVerifiers []PathVerifier `json:"PathVerifiers,omitempty"`
}
