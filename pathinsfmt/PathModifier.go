package pathinsfmt

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
)

type PathModifier struct {
	chmodins.BaseRwxInstructions
	ChmodCommands *ChmodCommands `json:"ChmodCommands,omitempty"`
	Chown         *Chown         `json:"Chown,omitempty"`
	ChangeGroup   *ChangeGroup   `json:"ChangeGroup,omitempty"`
}

func (p *PathModifier) HasChmodCommands() bool {
	return p != nil &&
		p.ChmodCommands != nil &&
		p.ChmodCommands.HasAnyItem()
}

func (p *PathModifier) HasChangeGroup() bool {
	return p != nil &&
		p.ChangeGroup != nil
}

func (p *PathModifier) HasChown() bool {
	return p != nil &&
		p.Chown != nil
}
