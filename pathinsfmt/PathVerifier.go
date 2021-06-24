package pathinsfmt

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type PathVerifier struct {
	BaseUserNamePlusGroupName
	chmodins.BaseRwxInstructions
}

func (p *PathVerifier) HasRwxInstructions() bool {
	return p != nil && p.RwxInstructions != nil && p.BaseRwxInstructions.HasAnyItem()
}
