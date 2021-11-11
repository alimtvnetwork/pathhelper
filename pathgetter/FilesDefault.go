package pathgetter

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func FilesDefault(
	isNormalize bool,
	exploringPaths ...string,
) *errstr.Results {
	if exploringPaths == nil {
		return errstr.Empty.Results()
	}

	return FilesUsingPathsPtr(
		isNormalize,
		osconsts.PathSeparator,
		exploringPaths)
}
