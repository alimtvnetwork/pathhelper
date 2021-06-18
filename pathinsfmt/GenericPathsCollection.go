package pathinsfmt

type GenericPathsCollection struct {
	lazyFlatPaths []string
	SimilarPaths  []SimilarPaths `json:"SimilarPaths,omitempty"`
	AllDiffPaths  []AllDiffPaths `json:"AllDiffPaths,omitempty"`
	DynamicPaths  *DynamicPaths  `json:"DynamicPaths,omitempty"`
}

// Length of len(receiver.SimilarPaths) + len(receiver.AllDiffPaths) + items in DynamicPaths (not all specific paths)
func (receiver *GenericPathsCollection) Length() int {
	length := len(receiver.SimilarPaths) + len(receiver.AllDiffPaths)

	if receiver.DynamicPaths == nil {
		return length
	}

	return length + receiver.DynamicPaths.Length()
}

func (receiver *GenericPathsCollection) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *GenericPathsCollection) HasAnyItem() bool {
	return receiver.Length() > 0
}
