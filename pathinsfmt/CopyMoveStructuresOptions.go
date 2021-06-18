package pathinsfmt

type CopyMoveStructuresOptions struct {
	BasePathModifiers
	BaseIsRename
	IsClearBeforeStart bool   `json:"IsClearBeforeStart"`
	IsMove             bool   `json:"IsMove"`
	PermissionOptions  string `json:"PermissionOptions,omitempty"`
	OverwriteConfig    string `json:"OverwriteConfig,omitempty"`
}
