package recursivepaths

import "gitlab.com/auk-go/errorwrapper/errdata/errstr"

func Directories(
	rootPath string,
) *errstr.Results {
	return DirectoriesOptions(
		true,
		false,
		false,
		rootPath)
}
