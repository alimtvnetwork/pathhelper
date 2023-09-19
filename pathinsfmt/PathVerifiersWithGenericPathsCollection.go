package pathinsfmt

import "gitlab.com/auk-go/core/coreinstruction"

type PathVerifiersWithGenericPathsCollection struct {
	coreinstruction.BaseSpecPlusRequestIds
	BaseGenericPathsCollection
	PathVerifiers *PathVerifiers `json:"PathVerifiers,omitempty"`
}

func (p *PathVerifiersWithGenericPathsCollection) IsEitherEmpty() bool {
	return p == nil ||
		p.PathVerifiers == nil ||
		p.BaseGenericPathsCollection.GenericPathsCollection == nil ||
		p.PathVerifiers.IsEmpty() ||
		p.BaseGenericPathsCollection.IsEmptyGenericPathsCollection()
}
