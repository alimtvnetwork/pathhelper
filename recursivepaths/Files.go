package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

// Files normalize, expand false
func Files(rootPath string) *errstr.Results {
	return FilesOptions(
		true,
		false,
		false,
		rootPath)
}
