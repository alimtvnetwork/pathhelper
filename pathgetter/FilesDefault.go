package pathgetter

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
)

func FilesDefault(
	isNormalize bool,
	location string,
) *errstr.Results {
	return Files(
		isNormalize,
		osconsts.PathSeparator,
		location)
}
