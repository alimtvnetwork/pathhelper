package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func ReturnAsIs(
	isNormalize bool,
	isExpandEnv bool,
	rootPath string,
) *errstr.Results {
	fixedPath := pathjoin.FixPath(
		isNormalize,
		isExpandEnv,
		rootPath)

	return errstr.New.Results.SpreadValuesOnly(
		fixedPath)
}
