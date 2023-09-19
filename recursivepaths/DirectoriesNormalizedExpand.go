package recursivepaths

import "gitlab.com/auk-go/errorwrapper/errdata/errstr"

func DirectoriesNormalizedExpand(
	rootPath string,
) *errstr.Results {
	return DirectoriesOptions(
		true,
		true,
		true,
		rootPath)
}
