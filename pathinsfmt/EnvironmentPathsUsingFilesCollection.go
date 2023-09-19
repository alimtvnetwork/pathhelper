package pathinsfmt

import "gitlab.com/auk-go/core/reqtype"

type EnvironmentPathsUsingFilesCollection struct {
	BaseLocationCollection
	ModifyAs reqtype.Request `json:"ModifyAs"`
}
