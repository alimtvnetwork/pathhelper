package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

// FilesNonRecursive normalize, expand false
func FilesNonRecursive(rootPath string) *errstr.Results {
	return FilesOptions(
		false,
		false,
		false,
		rootPath)
}
