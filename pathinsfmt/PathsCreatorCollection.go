package pathinsfmt

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type PathsCreatorCollection struct {
	PathsCreateInstruction  []BasePathsCreator `json:"PathsCreateInstruction,omitempty"`
	IsIgnoreOnExist         bool
	IsDeleteAllBeforeCreate bool
	ApplyRwx                *chmodins.RwxOwnerGroupOther
	ApplyUserGroup          *UserGroupName
}
