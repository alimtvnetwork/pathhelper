package pathinsfmt

type CopyPath struct {
	SourceDestinationPlusCompiled
	BaseIsRename
	Rename  string           `json:"Rename,omitempty"`
	Options *CopyPathOptions `json:"Options,omitempty"`
}

func (it *CopyPath) HasOptions() bool {
	return it != nil && it.Options != nil
}
