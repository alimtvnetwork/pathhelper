package pathinsfmt

type CopyMoveStructuresOptions struct {
	*BasePathModifiers
	*BaseIsRename
	IsClearBeforeStart bool   `json:"IsClearBeforeStart"`
	IsMove             bool   `json:"IsMove"`
	PermissionOptions  string `json:"PermissionOptions,omitempty"`
	OverwriteConfig    string `json:"OverwriteConfig,omitempty"`
}

type BaseSourceDestination struct {
	Source      string `json:"Source,omitempty"`
	Destination string `json:"Destination,omitempty"`
}

type CopyMoveStructure struct {
	BasePathModifiers
	*BaseSourceDestination
	Rename  string                     `json:"Rename,omitempty"`
	Options *CopyMoveStructuresOptions `json:"Options,omitempty"`
}
