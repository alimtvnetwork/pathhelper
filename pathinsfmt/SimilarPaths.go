package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type SimilarPaths struct {
	RootPath         string   `json:"RootPath"`
	RelativePaths    []string `json:"RelativePaths"`
	IsNormalizeApply bool     `json:"IsNormalizeApply"`
}

func (s *SimilarPaths) Length() int {
	return len(s.RelativePaths)
}

func (s *SimilarPaths) IsEmpty() bool {
	return s.Length() == 0
}

func (s *SimilarPaths) HasAnyItem() bool {
	return s.Length() > 0
}

func (s *SimilarPaths) FlatPaths(isIncludeRootAsClone bool) []string {
	if s.IsEmpty() && !isIncludeRootAsClone {
		return []string{}
	}

	slice := make(
		[]string,
		constants.Zero,
		s.Length()+constants.Capacity2)

	if isIncludeRootAsClone {
		slice = append(slice, s.RootPath)
	}

	root := s.RootPath

	for _, relativePath := range s.RelativePaths {
		joinedPath := pathjoin.JoinNormalizedIf(
			s.IsNormalizeApply,
			root,
			relativePath)

		slice = append(
			slice,
			joinedPath)
	}

	return slice
}
