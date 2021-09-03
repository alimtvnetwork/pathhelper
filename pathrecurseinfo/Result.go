package pathrecurseinfo

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

type Result struct {
	Root            string
	PathStat        *chmodhelper.PathExistStat
	IsInvalidResult bool
	IsRelative      bool
	PathsResult     *PathsResult
	ErrorWrapper    *errorwrapper.Wrapper
}

func (it *Result) IsEmptyPathStat() bool {
	if it == nil || it.PathStat == nil {
		return true
	}

	return it.PathStat != nil
}

func (it *Result) HasSafeItems() bool {
	return it.Length() > 0 &&
		it.PathsResult.IsExist &&
		!it.IsInvalidResult &&
		it.ErrorWrapper.HasError()
}

func (it *Result) Paths() []string {
	if it == nil || it.PathsResult == nil {
		return []string{}
	}

	return it.PathsResult.ExpandingPaths.Items
}

func (it *Result) PathsString() string {
	if it == nil || it.PathsResult == nil {
		return constants.EmptyString
	}

	return strings.Join(
		it.PathsResult.ExpandingPaths.Items,
		constants.NewLineUnix)
}

func (it *Result) HasIssuesOrEmpty() bool {
	return it.IsEmpty() ||
		it.ErrorWrapper.HasError() ||
		it.IsInvalidResult
}

func (it *Result) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Result) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Result) Length() int {
	if it == nil || it.PathsResult == nil {
		return 0
	}

	return it.PathsResult.ExpandingPaths.Length()
}

func (it *Result) Strings() []string {
	return it.StringsResults().ValueNonPtr()
}

func (it *Result) FilterStringsResults(filter pathfuncs.Filter) *errstr.Results {
	if it.HasIssuesOrEmpty() {
		return it.StringsResults()
	}

	sep := osconsts.PathSeparator
	slice := stringslice.MakeDefault(it.Length())
	var errSlice []string

	for _, fullPath := range it.PathsResult.ExpandingPaths.Items {
		// TODO : Fix error swallow, probably no need to fix,
		//        as error already checked and fix at first
		fileInfo, _ := os.Stat(fullPath)
		hasFileInfo := fileInfo != nil

		arg := &pathfuncs.FilterArg{
			RootPath:    it.Root,
			FileName:    splitinternal.GetFileNameWithExt(fullPath),
			FullPath:    fullPath,
			Separator:   sep,
			IsFile:      hasFileInfo && !fileInfo.IsDir(),
			IsDirectory: hasFileInfo && fileInfo.IsDir(),
			FileInfo:    fileInfo,
		}

		filterResult := filter(arg)

		if filterResult.ErrorWrapper.HasError() {
			errSlice = append(
				errSlice,
				filterResult.ErrorWrapper.FullString())
		}

		if filterResult.IsKeep {
			slice = append(slice, fullPath)
		}

		if filterResult.IsBreak {
			break
		}
	}

	err := msgtype.SliceToError(errSlice)

	if err != nil {
		return errstr.EmptyResultsWithError(
			errnew.Path(
				errtype.PathExpand,
				err,
				it.Root))
	}

	return errstr.EmptyErrorResults(
		slice...)
}

func (it *Result) FilterResults(filter pathfuncs.Filter) *Result {
	if it.HasIssuesOrEmpty() {
		return InvalidResult(
			it.Root,
			it.ErrorWrapper,
			it.PathStat)
	}

	results := it.FilterStringsResults(filter)

	return &Result{
		Root:            it.Root,
		PathStat:        chmodhelper.GetPathExistStat(it.Root),
		IsInvalidResult: it.IsInvalidResult,
		IsRelative:      it.IsRelative,
		PathsResult: &PathsResult{
			ExpandingPaths: results.SimpleSlice(),
			IsExist:        it.PathsResult.IsExist,
			IsFile:         it.PathsResult.IsFile,
			IsDir:          it.PathsResult.IsDir,
		},
		ErrorWrapper: results.ErrorWrapper,
	}
}

func (it *Result) StringsResults() *errstr.Results {
	if it == nil {
		return errstr.EmptyResultsWithError(
			errnew.NilOrEmpty)
	}

	if it.ErrorWrapper.HasError() || it.IsEmpty() {
		return errstr.EmptyResultsWithError(
			it.ErrorWrapper)
	}

	return errstr.NewResults(
		it.ErrorWrapper,
		it.PathsResult.ExpandingPaths.Items...)
}

func (it *Result) Clone() *Result {
	if it == nil {
		return nil
	}

	return &Result{
		Root:            it.Root,
		PathStat:        chmodhelper.GetPathExistStat(it.Root),
		IsInvalidResult: it.IsInvalidResult,
		IsRelative:      it.IsRelative,
		PathsResult:     it.PathsResult.Clone(),
		ErrorWrapper:    it.ErrorWrapper,
	}
}
