package pathrecurseinfo

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
)

func InvalidResult(
	root string,
	errorWrapper *errorwrapper.Wrapper,
	stat *chmodhelper.PathExistStat,
) *Result {
	return &Result{
		Root:            root,
		IsInvalidResult: true,
		PathStat:        stat,
		ErrorWrapper:    errorWrapper,
		PathsResult:     EmptyPathsResult(),
	}
}
