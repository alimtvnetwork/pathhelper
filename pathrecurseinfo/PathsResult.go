package pathrecurseinfo

import (
	"path"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type PathsResult struct {
	ExpandingPaths *corestr.SimpleSlice
	IsExist,
	IsFile,
	IsDir bool
}

func (it *PathsResult) JoinWithRoot(
	root string,
	isNormalize bool,
) *corestr.SimpleSlice {
	if it.ExpandingPaths.IsEmpty() {
		return corestr.NewSimpleSlice(0)
	}

	rootFix := normalize.PathUsingSingleIf(isNormalize, root)

	newSlice := make([]string, it.ExpandingPaths.Length())

	if isNormalize {
		for _, item := range it.ExpandingPaths.Items {
			newPath := pathjoin.JoinSimple(rootFix, item)

			newSlice = append(newSlice, newPath)
		}

		return corestr.NewSimpleSliceUsing(false, newSlice)
	}

	for _, item := range it.ExpandingPaths.Items {
		newPath := path.Join(rootFix, item)

		newSlice = append(newSlice, newPath)
	}

	return corestr.NewSimpleSliceUsing(
		false,
		newSlice)
}

func (it *PathsResult) Clone() *PathsResult {
	if it == nil {
		return nil
	}

	return &PathsResult{
		ExpandingPaths: corestr.NewSimpleSliceUsing(
			true,
			it.ExpandingPaths.Items),
		IsExist: it.IsExist,
		IsFile:  it.IsFile,
		IsDir:   it.IsDir,
	}
}

func (it *PathsResult) ConcatNew(
	isClone bool,
	other *PathsResult,
) *PathsResult {
	if other == nil && !isClone {
		return it
	}

	if other == nil && isClone {
		return it.Clone()
	}

	slice := it.ExpandingPaths.ConcatNewSimpleSlices(
		other.ExpandingPaths)

	return &PathsResult{
		ExpandingPaths: slice,
		IsExist:        it.IsExist && other.IsExist,
		IsFile:         it.IsFile && other.IsFile,
		IsDir:          it.IsDir && other.IsDir,
	}
}
