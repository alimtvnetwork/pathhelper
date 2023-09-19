package pathinsfmt

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
)

type PathVerifier struct {
	UserGroupName
	chmodins.BaseRwxInstructions
}

func (p *PathVerifier) HasRwxInstructions() bool {
	return p != nil && p.RwxInstructions != nil && p.BaseRwxInstructions.HasAnyItem()
}
