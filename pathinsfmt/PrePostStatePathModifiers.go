package pathinsfmt

type PrePostStatePathModifiers struct {
	PrePathModifiers  []PathVerifiersWithGenericPathsCollection `json:"PrePathModifiers,omitempty"`
	PostPathModifiers []PathVerifiersWithGenericPathsCollection `json:"PostPathModifiers,omitempty"`
}
