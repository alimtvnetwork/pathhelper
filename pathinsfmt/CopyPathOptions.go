package pathinsfmt

type CopyPathOptions struct {
	IsClearAtFirst    bool          `json:"IsClearAtFirst,omitempty"` // removes all before the action
	IsOverwrite       bool          `json:"IsOverwrite,omitempty"`
	IsCopyRwx         bool          `json:"IsCopyRwx,omitempty"`
	IsMove            bool          `json:"IsMove,omitempty"`
	ApplyPathModifier *PathModifier `json:"ApplyPathModifier,omitempty"`
}
