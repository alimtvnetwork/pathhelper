package pathinsfmt

import "gitlab.com/evatix-go/core/reqtype"

type EnvironmentVariable struct {
	Name  string `json:"Name,omitempty"`
	Value string `json:"Value,omitempty"`
}

type EnvironmentPaths struct {
	ModifyAs reqtype.Request `json:"ModifyAs"`
	Paths    []string        `json:"Locations,omitempty"`
}

type EnvironmentPathsUsingGenericPaths struct {
	ModifyAs     reqtype.Request         `json:"ModifyAs"`
	GenericPaths *GenericPathsCollection `json:"GenericPathsCollection,omitempty"`
}

type EnvironmentPathsUsingFilesCollection struct {
	ModifyAs        reqtype.Request     `json:"ModifyAs"`
	FilesCollection *LocationCollection `json:"LocationCollection,omitempty"`
}
