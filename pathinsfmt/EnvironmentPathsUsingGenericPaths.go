package pathinsfmt

import (
	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/reqtype"
)

type EnvironmentPathsUsingGenericPaths struct {
	coreinstruction.BaseSpecPlusRequestIds
	BaseGenericPathsCollection
	ModifyAs reqtype.Request `json:"ModifyAs"`
}
