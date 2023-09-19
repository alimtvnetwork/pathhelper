package pathjoin

import (
	"path"
	"path/filepath"
	"runtime"

	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func JoinWithCurDirUsingCallStack(
	isNormalize,
	isExpandEnv bool,
	frameStackSkip int,
	paths ...string,
) string {
	_, b, _, _ := runtime.Caller(frameStackSkip)
	finalSlice := stringslice.PrependLineNew(filepath.Dir(b), paths)
	joined := path.Join(finalSlice...)

	expanded := expandpath.ExpandVariablesIf(
		isExpandEnv,
		joined)

	return normalize.PathUsingSeparatorIf(
		false,
		isNormalize,
		isNormalize,
		osconsts.PathSeparator,
		expanded)
}
