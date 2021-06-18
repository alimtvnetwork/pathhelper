package pathinsfmt

type FilesSelector struct {
	Path        string      `json:"Path"`
	Filters     []string    `json:"Filters,omitempty"`
	SkipFilters []string    `json:"SkipFilters,omitempty"`
	Extensions  []string    `json:"Extensions,omitempty"`
	Processors  []string    `json:"Processors,omitempty"`
	Attributes  *Attributes `json:"Attributes,omitempty"`
}
