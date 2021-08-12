package elitepath

type Directory struct {
	WithPermission
	IsClearDir                bool `json:"IsClearDir,omitempty"`
	IsApplyDefaultChmodGroups bool `json:"IsApplyDefaultChmodGroups,omitempty"`
}
