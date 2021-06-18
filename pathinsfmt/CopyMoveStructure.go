package pathinsfmt

type CopyMoveStructure struct {
	BasePathModifiers
	BaseSourceDestination
	Rename  *string                    `json:"Rename,omitempty"`
	Options *CopyMoveStructuresOptions `json:"Options,omitempty"`
}
