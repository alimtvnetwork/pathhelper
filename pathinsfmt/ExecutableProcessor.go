package pathinsfmt

type ExecutableProcessor struct {
	Name            string   `json:"Name"`
	IsEnabled       bool     `json:"IsEnabled,omitempty"`
	IsNormalizePath bool     `json:"IsNormalizePath,omitempty"`
	ProcessorPath   string   `json:"ProcessorPath,omitempty"`
	BinaryPath      string   `json:"BinaryPath"`
	Args            []string `json:"Args,omitempty"`
}
