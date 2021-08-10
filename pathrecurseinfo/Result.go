package pathrecurseinfo

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errinf"
)

type Result struct {
	Root            string
	PathStat        *chmodhelper.PathExistStat
	IsInvalidResult bool
	IsRelative      bool
	PathsResult     *PathsResult
	errinf.ErrWrapper
}

func (it *Result) IsEmptyPathStat() bool {
	if it == nil || it.PathStat == nil {
		return true
	}

	return it.PathStat != nil
}

func (it *Result) IsExist() bool {
	if it == nil || it.PathsResult == nil {
		return false
	}

	return it.PathsResult.IsExist
}

func (it *Result) Length() int {
	if it == nil || it.PathsResult == nil {
		return 0
	}

	return it.PathsResult.ExpandingPaths.Length()
}
