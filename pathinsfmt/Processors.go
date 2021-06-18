package pathinsfmt

type Processors struct {
	Name          string   `json:"Name"`
	IsEnabled     bool     `json:"IsEnabled"`
	ProcessorPath string   `json:"ProcessorPath"`
	Exe           string   `json:"Exe"`
	Args          []string `json:"Args,omitempty"`
}
