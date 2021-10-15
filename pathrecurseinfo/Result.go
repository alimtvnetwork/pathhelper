package pathrecurseinfo

import (
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
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
	return it.StringsResults().SafeValues()
}

func (it *Result) FilterFullPaths(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	filter pathfuncs.Filter,
) *corestr.SimpleSlice {
	if it == nil || it.PathsResult == nil || it.PathsResult.ExpandingPaths.IsEmpty() {
		return corestr.EmptySimpleSlice()
	}

	hasErr := it.ErrorWrapper.HasError()
	if hasErr {
		errCollection.AddWrapperPtr(it.ErrorWrapper)
	}

	if !isContinueOnError && hasErr {
		return corestr.EmptySimpleSlice()
	}

	return pathfuncs.FilterFullPaths(
		false,
		errCollection,
		filter,
		it.Root,
		it.PathsResult.ExpandingPaths.Items...)
}

func (it *Result) SimpleFilterFullPathsAsync(
	filter pathfuncs.SimpleFilter,
) *errstr.Results {
	if it.HasIssuesOrEmpty() {
		return it.EmptyErrorResults()
	}

	return pathfuncs.SimpleFilterFullPathsAsync(
		false,
		filter,
		it.PathsResult.ExpandingPaths.Items...)
}

func (it *Result) FilterResults(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	filter pathfuncs.Filter,
) *Result {
	if it.HasIssuesOrEmpty() {
		return InvalidResult(
			it.Root,
			it.ErrorWrapper,
			it.PathStat)
	}

	stateTracker := errCollection.StateTracker()
	results := it.FilterFullPaths(
		isContinueOnError,
		errCollection,
		filter)

	errWp := errnew.EmptyPtr

	if stateTracker.IsFailed() {
		errWp = errnew.MessagesPtr(
			errtype.AlreadyDefined,
			"errors are already collected in the error collection.")
	}

	return &Result{
		Root:            it.Root,
		PathStat:        it.PathStat,
		IsInvalidResult: it.IsInvalidResult,
		IsRelative:      it.IsRelative,
		PathsResult: &PathsResult{
			ExpandingPaths: results,
			IsExist:        it.PathsResult.IsExist,
			IsFile:         it.PathsResult.IsFile,
			IsDir:          it.PathsResult.IsDir,
		},
		ErrorWrapper: errWp,
	}
}

func (it *Result) StringsResults() *errstr.Results {
	if it == nil {
		return errstr.EmptyResultsWithError(
			errnew.NullSimple(it))
	}

	if it.ErrorWrapper.HasError() || it.IsEmpty() {
		return errstr.EmptyResultsWithError(
			it.ErrorWrapper)
	}

	return errstr.NewResults(
		it.ErrorWrapper,
		it.PathsResult.ExpandingPaths.Items...)
}

func (it *Result) EmptyErrorResults() *errstr.Results {
	if it == nil {
		return errstr.EmptyResultsWithError(
			errnew.NullSimple(it))
	}

	if it.ErrorWrapper.HasError() || it.IsEmpty() {
		return errstr.EmptyResultsWithError(
			it.ErrorWrapper)
	}

	return errstr.EmptyResults()
}

func (it *Result) Clone(isDeepClone bool) *Result {
	if it == nil {
		return nil
	}

	return &Result{
		Root:            it.Root,
		PathStat:        chmodhelper.GetPathExistStat(it.Root),
		IsInvalidResult: it.IsInvalidResult,
		IsRelative:      it.IsRelative,
		PathsResult:     it.PathsResult.Clone(isDeepClone),
		ErrorWrapper:    it.ErrorWrapper,
	}
}

func (it *Result) Dispose() {
	if it == nil {
		return
	}

	it.Root = constants.EmptyString
	it.PathStat = nil
	it.ErrorWrapper.Dispose()
	it.PathsResult.Dispose()
}
