package pathinsfmt

type CliRunner struct {
	FilesSelector []FilesSelector `json:"FilesSelector,omitempty"`
	Processors    []Processors    `json:"Processors,omitempty"`
}
