package pathinsfmt

import "gitlab.com/evatix-go/core/coreinstruction"

type PathVerifiersWithGenericPathsCollection struct {
	coreinstruction.BaseSpecPlusRequestIds
	BaseGenericPathsCollection
	PathVerifiers []PathVerifier `json:"PathVerifiers,omitempty"`
}
