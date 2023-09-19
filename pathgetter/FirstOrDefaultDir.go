package pathgetter

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func FirstOrDefaultDir(
	isNormalize bool,
	rootPath string,
) *errstr.Result {
	results := Dirs(
		osconsts.PathSeparator,
		rootPath,
		isNormalize)

	return results.FirstOrDefaultResult()
}
