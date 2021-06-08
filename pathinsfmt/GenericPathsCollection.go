package pathinsfmt

type SimilarPaths struct {
	RootPath         string    `json:"RootPath"`
	RelativePaths    *[]string `json:"RelativePaths"`
	IsNormalizeApply bool      `json:"IsNormalizeApply"`
}

type AllDiffPaths struct {
	Paths            *[]string `json:"Paths"`
	IsNormalizeApply bool      `json:"IsNormalizeApply"`
}

type DynamicPaths struct {
	Vars         *[]PathVar      `json:"Vars"`
	AllDiffPaths *[]AllDiffPaths `json:"AllDiffPaths"`
}

type GenericPathsCollection struct {
	SimilarPaths *[]SimilarPaths `json:"SimilarPaths"`
	AllDiffPaths *[]AllDiffPaths `json:"AllDiffPaths"`
	DynamicPaths *DynamicPaths   `json:"DynamicPaths"`
}

type BaseGenericPathsCollection struct {
	GenericPathsCollection *GenericPathsCollection `json:"GenericPathsCollection,omitempty"`
}
