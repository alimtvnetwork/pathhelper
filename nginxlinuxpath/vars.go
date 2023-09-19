package nginxlinuxpath

import "gitlab.com/auk-go/core/coredata/corestr"

var (
	DefaultDirStructure = GetFullDirStructure(
		true,
		DefaultDirChmod,
		DefaultRoot)
	defaultMimeTypesPath = corestr.SimpleStringOnce{}
)
