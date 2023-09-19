package pathgetter

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
)

func FilterDefault(
	isNormalize bool,
	isIgnoreOnError bool,
	rootPath string,
	filter pathfuncs.Filter,
) []*pathfuncs.FilterResult {
	return Filter(
		osconsts.PathSeparator,
		rootPath,
		isNormalize,
		isIgnoreOnError,
		filter)
}
