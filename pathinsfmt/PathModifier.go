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
