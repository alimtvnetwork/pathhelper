package ds

type NestedPathsParam struct {
	IsContinueOnError bool
	results           *NestedPathsResults
	Separator         string
}

func (nestedPathsParam *NestedPathsParam) Results() *NestedPathsResults {
	if nestedPathsParam.results != nil {
		return nestedPathsParam.results
	}

	if nestedPathsParam.results == nil {
		nestedPathsParam.results =
			NewNestedPathsResults(nestedPathsParam)
	}

	return nestedPathsParam.results
}
