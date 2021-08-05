package pathinsfmt

import "os"

// Download Use aria2c
// Reference : https://aria2.github.io/manual/en/html/aria2c.html#options
type Download struct {
	URL              string      `json:"URL,omitempty"`
	Destination      string      `json:"Destination,omitempty"`
	FileName         string      `json:"FileName,omitempty"`
	ParallelRequests byte        `json:"ParallelRequests,omitempty"`
	MaxRetries       byte        `json:"MaxRetries,omitempty"`
	IsCreateDir      bool        `json:"IsCreateDir,omitempty"`
	IsClearDir       bool        `json:"IsClearDir,omitempty"`
	FileModeDir      os.FileMode `json:"FileModeDir,omitempty"`
}
