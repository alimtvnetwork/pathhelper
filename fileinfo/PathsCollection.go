package fileinfo

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
)

type PathsCollection struct {
	rootPath       string
	pathWrappers   *[]PathWrapper
	allPaths       *[]string
	directories    *[]string
	files          *[]string
	Error          *errorwrapper.Wrapper
	parentWrappers *Wrappers
}

func (pathsCollection *PathsCollection) AllPaths(separator string, isContinueOnError bool) (
	allPaths *[]string,
	errorWrappersCollection *errwrappers.Collection,
) {
	if pathsCollection.allPaths != nil {
		empty := errwrappers.Empty()

		return pathsCollection.allPaths, empty
	}

	allPaths, errWrappers := recursiveinternal.GetPaths(
		separator,
		pathsCollection.rootPath,
		isContinueOnError)

	if errWrappers.IsEmpty() {
		pathsCollection.allPaths = allPaths

		return allPaths, errWrappers
	}

	if errWrappers.IsEmpty() {
		pathsCollection.allPaths = allPaths
	}

	return allPaths, errWrappers
}

func NewPaths(rootPath string, capacity int) *PathsCollection {
	paths := make([]PathWrapper, 0, capacity)

	return &PathsCollection{
		rootPath:     rootPath,
		pathWrappers: &paths,
	}
}

func NewPathsUsingWrappers(rootPath string, wrappers *Wrappers) *PathsCollection {
	if wrappers == nil {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   nil,
			Error:          errnew.EmptyPtr,
			parentWrappers: wrappers,
		}
	}

	if wrappers.IsEmpty() {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   nil,
			Error:          wrappers.Error,
			parentWrappers: wrappers,
		}
	}

	paths := make([]PathWrapper, wrappers.Length())
	for i, wrapper := range *wrappers.collection {
		paths[i] = PathWrapper{
			Path:        wrapper.RawPath,
			IsDirectory: wrapper.IsDirectory,
		}
	}

	return &PathsCollection{
		rootPath:       rootPath,
		pathWrappers:   &paths,
		Error:          wrappers.Error,
		parentWrappers: wrappers,
	}
}

func NewPathsUsing(directoryPath string) *PathsCollection {
	wrappers := NewWrappersPtr(directoryPath)

	return NewPathsUsingWrappers(directoryPath, wrappers)
}

func NewPathsUsingPaths(rootPath string, paths *[]string) *PathsCollection {
	if paths == nil {
		return NewPaths(
			rootPath,
			0)
	}

	wrappers := NewPaths(rootPath, len(*paths))
	wrappers.allPaths = paths

	return wrappers
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
		pathsCollection.Error.HasError() ||
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

func (pathsCollection *PathsCollection) Add(wrapper PathWrapper) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		wrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) AddPtr(wrapper *PathWrapper) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		*wrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) IsContains(
	path string,
	isCaseSensitive bool,
) bool {
	return isstrsinternal.ContainsPtrSimple(
		pathsCollection.PathsAsStrings(),
		path,
		0,
		isCaseSensitive)
}

func (pathsCollection *PathsCollection) AddWrapper(pathWrapper PathWrapper) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		pathWrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) PathsAsStrings() *[]string {
	if pathsCollection.allPaths != nil {
		return pathsCollection.allPaths
	}

	collection := make([]string, pathsCollection.Length())

	for i, pathWrapper := range *pathsCollection.pathWrappers {
		collection[i] = pathWrapper.Path
	}

	pathsCollection.allPaths = &collection

	return pathsCollection.allPaths
}
