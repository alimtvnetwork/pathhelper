package pathinsfmt

type PrePostStatePathModifiers struct {
	PrePathModifiers  []PathModifier `json:"PrePathModifiers,omitempty"`
	PostPathModifiers []PathModifier `json:"PostPathModifiers,omitempty"`
}
