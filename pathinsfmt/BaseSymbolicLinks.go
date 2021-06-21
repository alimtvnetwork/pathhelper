package pathinsfmt

import "gitlab.com/evatix-go/core/coreinstruction"

type BaseSymbolicLinks struct {
	coreinstruction.BaseSpecPlusRequestIds
	SymbolicLinks []SymbolicLink `json:"SymbolicLinks,omitempty"`
}
