package pathinsfmt

type PrePostStatePathModifiers struct {
	Pre  BasePathsWithVerifiers `json:"PrePathModifiers,omitempty"`
	Post BasePathsWithVerifiers `json:"PostPathModifiers,omitempty"`
}
