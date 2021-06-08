package pathinsfmt

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type EnvironmentVariable struct {
	Name  string `json:"Name,omitempty"`
	Value string `json:"Value,omitempty"`
}

type SymbolicLink struct {
	Src           string `json:"Src"`
	Dst           string `json:"Dst"`
	IsForce       bool   `json:"IsForce"`
	IsSkipOnError bool   `json:"IsSkipOnError"`
}

type BaseSymbolicLinks struct {
	SymbolicLinks *[]SymbolicLink `json:"SymbolicLinks,omitempty"`
}

type ChmodCommand struct {
	BaseIsRecursive
	Command string `json:"Command,omitempty"`
}

type BaseGroupName struct {
	GroupName *string `json:"GroupName,omitempty"` // Not define or empty string or * means keeping the existing one
}

type BaseUserNamePlusGroupName struct {
	*BaseGroupName
	UserName *string `json:"UserName,omitempty"` // Not define or empty string or * means keeping the existing one
}

type BaseIsRecursive struct {
	IsRecursive bool `json:"IsRecursive"`
}

type Chown struct {
	BaseIsRecursive
	BaseUserNamePlusGroupName
}

type ChangeGroup struct {
	BaseIsRecursive
	GroupName string `json:"GroupName,omitempty"`
}

type PathModifier struct {
	*chmodins.BaseRwxInstructions
	GenericPathsCollection *GenericPathsCollection `json:"GenericPathsCollection,omitempty"`
	ChmodCommand           *ChmodCommand           `json:"ChmodCommand,omitempty"`
	Chown                  *Chown                  `json:"Chown,omitempty"`
	ChangeGroup            *ChangeGroup            `json:"ChangeGroup,omitempty"`
}

type BasePathModifiers struct {
	PathModifiers *[]PathModifier `json:"PathModifiers,omitempty"`
}

type PrePostStatePathModifiers struct {
	PrePathModifiers  *[]PathModifier `json:"PrePathModifiers,omitempty"`
	PostPathModifiers *[]PathModifier `json:"PostPathModifiers,omitempty"`
}

type PathVerifier struct {
	*BaseUserNamePlusGroupName
	*chmodins.BaseRwxInstructions
	Path string `json:"Path,omitempty"`
}

type BasePathVerifiers struct {
	PathVerifiers *[]PathVerifier `json:"PathVerifiers,omitempty"`
}

type BaseEnvironmentVariables struct {
	EnvVars *[]EnvironmentVariable `json:"EnvVars,omitempty"`
}

type BaseEnvPaths struct {
	EnvPaths *[]string `json:"EnvPaths,omitempty"`
}
