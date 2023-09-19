package fileinfo

import "gitlab.com/auk-go/errorwrapper"

type PathsCollectionDataModel struct {
	RootPath       string
	PathWrappers   *[]*SimplePathWrapper `json:"SimplePathWrappers"`
	Separator      string
	ErrorWrapper   *errorwrapper.Wrapper
	ParentWrappers *Wrappers
}
