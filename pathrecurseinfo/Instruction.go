package pathrecurseinfo

import (
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper/errinf"
)

type Instruction struct {
	Root                string
	ExcludingRootNames  []string //  path names contains in this will be ignored
	IsIncludeFilesOnly, //  includes only files if IsIncludeAll false
	IsRelativePath, // remove root path from paths
	IsIncludeDirsOnly, // includes only dir if IsIncludeAll false
	IsIncludeAll, // includes dir, files all
	IsExcludeRoot, // Don't include root path
	IsRecursive, // Recursively get paths if dir
	IsNormalize bool
	excludingNamesHashset *corestr.Hashset
}

func (it *Instruction) Result() *Result {
	return GetInstructionResult(it)
}

func (it *Instruction) HasAnyExcludingCondition() bool {
	return len(it.ExcludingRootNames) > 0
}

func (it *Instruction) ExcludingNamesHashset() *corestr.Hashset {
	if it.excludingNamesHashset != nil {
		return it.excludingNamesHashset
	}

	slicePtr := stringslice.SlicePtr(it.ExcludingRootNames)
	it.excludingNamesHashset = corestr.NewHashsetUsingStrings(
		slicePtr)

	return it.excludingNamesHashset
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

func InvalidResult(
	root string,
	errorWrapper errinf.ErrWrapper,
	stat *chmodhelper.PathExistStat,
) *Result {
	return &Result{
		Root:            root,
		IsInvalidResult: true,
		PathStat:        stat,
		ErrWrapper:      errorWrapper,
		PathsResult:     nil,
	}
}
