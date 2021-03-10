package fileinfo

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
)

type PathsCollection struct {
	rootPath          string
	pathWrappers      *[]*SimplePathWrapper
	allRecursivePaths *errstr.ResultsWithErrorCollection
	allRecursiveFiles *errstr.ResultsWithErrorCollection
	allRecursiveDirs  *errstr.ResultsWithErrorCollection
	directories       *[]string
	files             *[]string
	separator         string
	ErrorWrapper      *errorwrapper.Wrapper
	parentWrappers    *Wrappers
}

func NewPaths(rootPath, separator string, capacity int) *PathsCollection {
	paths := make([]*SimplePathWrapper, 0, capacity)

	return &PathsCollection{
		rootPath:     rootPath,
		pathWrappers: &paths,
		separator:    separator,
	}
}

func NewPathsUsingWrappers(
	rootPath, separator string,
	wrappers *Wrappers,
) *PathsCollection {
	if wrappers == nil {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   &[]*SimplePathWrapper{},
			ErrorWrapper:   errnew.EmptyPtr,
			parentWrappers: wrappers,
			separator:      separator,
		}
	}

	if wrappers.IsEmpty() {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   &[]*SimplePathWrapper{},
			ErrorWrapper:   wrappers.ErrorWrapper,
			parentWrappers: wrappers,
			separator:      separator,
		}
	}

	paths := make(
		[]*SimplePathWrapper,
		wrappers.Length())

	for i, wrapper := range *wrappers.Items {
		paths[i] = &SimplePathWrapper{
			Path:        wrapper.RawPath,
			IsDirectory: wrapper.IsDirectory,
		}
	}

	return &PathsCollection{
		rootPath:       rootPath,
		pathWrappers:   &paths,
		ErrorWrapper:   wrappers.ErrorWrapper,
		parentWrappers: wrappers,
		separator:      separator,
	}
}

func NewPathsUsing(
	directoryPath, separator string,
	isNormalize bool,
) *PathsCollection {
	wrappers := NewWrappersPtr(
		directoryPath,
		separator,
		isNormalize)

	return NewPathsUsingWrappers(
		directoryPath,
		separator,
		wrappers)
}

func NewPathsUsingPaths(
	rootPath, separator string,
	recursivePaths *[]string,
) *PathsCollection {
	if recursivePaths == nil {
		return NewPaths(
			rootPath,
			separator,
			0)
	}

	wrappers :=
		NewPaths(
			rootPath,
			separator,
			len(*recursivePaths))

	wrappers.allRecursivePaths =
		&errstr.ResultsWithErrorCollection{
			Values:        recursivePaths,
			ErrorWrappers: errwrappers.Empty(),
		}

	return wrappers
}

func (pathsCollection *PathsCollection) AllRecursivePaths() *errstr.ResultsWithErrorCollection {
	if pathsCollection.allRecursivePaths != nil {
		return pathsCollection.allRecursivePaths
	}

	allPaths, errWrappers := recursiveinternal.GetPaths(
		pathsCollection.separator,
		pathsCollection.rootPath,
		false)

	if errWrappers.IsEmpty() {
		pathsCollection.allRecursivePaths =
			&errstr.ResultsWithErrorCollection{
				Values:        allPaths,
				ErrorWrappers: errWrappers,
			}

		return pathsCollection.allRecursivePaths
	}

	pathsCollection.allRecursivePaths =
		&errstr.ResultsWithErrorCollection{
			Values:        core.EmptyStringsPtr(),
			ErrorWrappers: errWrappers,
		}

	return pathsCollection.allRecursivePaths
}

func (pathsCollection *PathsCollection) AllRecursiveFiles() *errstr.ResultsWithErrorCollection {
	if pathsCollection.allRecursiveFiles != nil {
		return pathsCollection.allRecursiveFiles
	}

	pathsCollection.allRecursiveFiles = recursiveinternal.GetFilesPaths(
		pathsCollection.separator,
		pathsCollection.rootPath,
		false)

	return pathsCollection.allRecursiveFiles
}

func (pathsCollection *PathsCollection) AllRecursiveDirs() *errstr.ResultsWithErrorCollection {
	if pathsCollection.allRecursiveDirs != nil {
		return pathsCollection.allRecursiveDirs
	}

	allPaths, errWrappers := recursiveinternal.GetDirectoryPaths(
		pathsCollection.separator,
		pathsCollection.rootPath,
		false)

	if errWrappers.IsEmpty() {
		pathsCollection.allRecursiveDirs = &errstr.ResultsWithErrorCollection{
			Values:        allPaths,
			ErrorWrappers: errWrappers,
		}

		return pathsCollection.allRecursiveDirs
	}

	pathsCollection.allRecursiveDirs = &errstr.ResultsWithErrorCollection{
		Values:        core.EmptyStringsPtr(),
		ErrorWrappers: errWrappers,
	}

	return pathsCollection.allRecursiveDirs
}

func (pathsCollection *PathsCollection) Directories() *[]string {
	if pathsCollection.directories != nil {
		return pathsCollection.directories
	}

	if pathsCollection.IsEmpty() {
		pathsCollection.directories =
			core.EmptyStringsPtr()

		return pathsCollection.directories
	}

	directories := make([]string, 0, pathsCollection.Length())

	for _, pathWrapper := range *pathsCollection.pathWrappers {
		if !pathWrapper.IsDirectory {
			continue
		}

		directories = append(directories, pathWrapper.Path)
	}

	pathsCollection.directories = &directories

	return pathsCollection.directories
}

func (pathsCollection *PathsCollection) Files() *[]string {
	if pathsCollection.files != nil {
		return pathsCollection.files
	}

	if pathsCollection.IsEmpty() {
		pathsCollection.files = core.EmptyStringsPtr()

		return pathsCollection.files
	}

	files := make([]string, 0, pathsCollection.Length())

	for _, pathWrapper := range *pathsCollection.pathWrappers {
		if pathWrapper.IsDirectory {
			continue
		}

		files = append(files, pathWrapper.Path)
	}

	pathsCollection.files = &files

	return pathsCollection.files
}

func (pathsCollection *PathsCollection) IsEmpty() bool {
	return pathsCollection.pathWrappers == nil ||
		pathsCollection.ErrorWrapper.HasError() ||
		len(*pathsCollection.pathWrappers) == 0
}

func (pathsCollection *PathsCollection) Length() int {
	if pathsCollection.pathWrappers == nil || *pathsCollection.pathWrappers == nil {
		return 0
	}

	return len(*pathsCollection.pathWrappers)
}

func (pathsCollection *PathsCollection) IsParentWrappersEmpty() bool {
	return pathsCollection.parentWrappers == nil ||
		pathsCollection.parentWrappers.IsEmpty()
}

func (pathsCollection *PathsCollection) HasParentWrappers() bool {
	return pathsCollection.parentWrappers != nil
}

func (pathsCollection *PathsCollection) ParentWrappers() *Wrappers {
	return pathsCollection.parentWrappers
}

func (pathsCollection *PathsCollection) Add(
	wrapper *SimplePathWrapper,
) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		wrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) AddPtr(
	wrapper *SimplePathWrapper,
) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		wrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) IsContains(
	path string,
	isCaseSensitive bool,
) bool {
	return isstrsinternal.ContainsPtrSimple(
		pathsCollection.AllRecursivePaths().Values,
		path,
		0,
		isCaseSensitive)
}

func (pathsCollection *PathsCollection) AddWrapper(
	pathWrapper *SimplePathWrapper,
) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		pathWrapper)

	return pathsCollection
}
