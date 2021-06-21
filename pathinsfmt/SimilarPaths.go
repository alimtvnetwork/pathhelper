package pathinsfmt

import (
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

func (s *SimilarPaths) FlatPaths() []string {
	if s.IsEmpty() {
		return []string{}
	}

	slice := make(
		[]string,
		s.Length())

	root := s.RootPath

	for i, relativePath := range s.RelativePaths {
		joinedPath := pathjoin.JoinNormalizedIf(
			s.IsNormalizeApply,
			root,
			relativePath)

		slice[i] = joinedPath
	}

	return slice
}
