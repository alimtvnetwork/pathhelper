package pathjoin

import (
	"path"

	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func JoinWithCurDir(
	isNormalize,
	isLongPathFix,
	isExpandEnv bool,
	paths ...string,
) string {
	finalSlice := stringslice.PrependLineNew(CurrentDirectory(), paths)
	joined := path.Join(finalSlice...)

	expanded := expandpath.ExpandVariablesIf(
		isExpandEnv,
		joined)

	return normalize.PathUsingSeparatorIf(
		isLongPathFix,
		isLongPathFix,
		isNormalize,
		osconsts.PathSeparator,
		expanded)
}
