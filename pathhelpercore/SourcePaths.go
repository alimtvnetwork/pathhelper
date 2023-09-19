package pathhelpercore

import "gitlab.com/auk-go/core/coredata/corestr"

type SourcePaths struct {
	RootDir     string              `json:"RootDir"`
	IsNormalize bool                `json:"IsNormalize"`
	Sources     *corestr.Collection `json:"Sources,omitempty"`
}
