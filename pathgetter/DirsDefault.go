package pathgetter

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func DirsDefault(
	isNormalize bool,
	rootPath string,
) *errstr.Results {
	return Dirs(
		osconsts.PathSeparator,
		rootPath,
		isNormalize)
}
