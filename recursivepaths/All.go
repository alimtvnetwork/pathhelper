package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func All(rootPath string) *errstr.Results {
	return AllOptions(
		false,
		false,
		rootPath)
}
