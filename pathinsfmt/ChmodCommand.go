package pathinsfmt

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type ChmodCommand struct {
	Condition *chmodins.Condition `json:"Condition,omitempty"`
	Commands  []string            `json:"Commands,omitempty"`
}
