package datamodel

type CliConfig struct {
	CliRunner *CliRunner `json:"CliRunner"`
}

type Duration struct {
	Value      int    `json:"Value"`
	Quantifier string `json:"Quantifier"`
}

type CachesRefresh struct {
	Duration Duration `json:"Duration"`
}

type Attributes struct {
	IsRecursive    bool          `json:"IsRecursive"`
	IsCache        bool          `json:"IsCache"`
	IsRedis        bool          `json:"IsRedis"`
	IsWriteToFiles bool          `json:"IsWriteToFiles"`
	CacheFilePath  string        `json:"CacheFilePath"`
	CachesRefresh  CachesRefresh `json:"CachesRefresh"`
}

type FilesSelector struct {
	Path        string     `json:"Path"`
	Filters     []string   `json:"Filters"`
	SkipFilters []string   `json:"SkipFilters"`
	Extensions  []string   `json:"Extensions"`
	Processors  []string   `json:"Processors"`
	Attribues   Attributes `json:"Attributes"`
}

type Processors struct {
	Name          string   `json:"Name"`
	IsEnabled     bool     `json:"IsEnabled"`
	ProcessorPath string   `json:"ProcessorPath"`
	Exe           string   `json:"Exe"`
	Args          []string `json:"Args"`
}

type CliRunner struct {
	FilesSelector []FilesSelector `json:"FilesSelector"`
	Processors    []Processors    `json:"Processors"`
}
