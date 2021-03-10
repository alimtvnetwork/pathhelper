package fileinfo

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

type Wrappers struct {
	RootPath            string
	separator           string
	collection          *[]Wrapper
	directories         *Wrappers
	files               *Wrappers
	Error               *errorwrapper.Wrapper
	pathsCollection     *PathsCollection
	recursiveDirs       *PathsCollection
	fileNamesCollection *FileNamesCollection
}

func (wrappers *Wrappers) RecursiveDirs() *PathsCollection {
	if wrappers.recursiveDirs != nil {
		return wrappers.recursiveDirs
	}

	allPaths, errorCollection :=
		recursiveinternal.GetDirectoryPaths(
			wrappers.separator,
			wrappers.RootPath,
			isErrorContinueDefault)

	errorCollection.HandleError()

	wrappers.recursiveDirs = NewPathsUsingPaths(
		wrappers.RootPath,
		wrappers.separator,
		allPaths)

	wrappers.recursiveDirs.parentWrappers = wrappers
	wrappers.recursiveDirs.directories = allPaths

	return wrappers.recursiveDirs
}

func (wrappers *Wrappers) RecursivePathsFilter(
	filter pathfuncs.Filter,
) *PathsCollection {
	allPaths, errorCollection := recursiveinternal.GetFilterPaths(
		wrappers.separator,
		wrappers.RootPath,
		isErrorContinueDefault,
		filter)

	errorCollection.HandleError()

	newPathsCollection := NewPathsUsingPaths(
		wrappers.RootPath,
		wrappers.separator,
		allPaths)

	newPathsCollection.parentWrappers = wrappers

	return newPathsCollection
}

func (wrappers *Wrappers) HasAny() bool {
	return !wrappers.IsEmpty()
}

func (wrappers *Wrappers) RootFiles() *Wrappers {
	if wrappers.files != nil {
		return wrappers.files
	}

	if wrappers.IsEmpty() {
		wrappers.files = EmptyWrappers()
	}

	files := make(
		[]Wrapper,
		0,
		wrappers.Length())

	for _, wrapper := range *wrappers.collection {
		if !wrapper.IsFile {
			continue
		}

		files = append(files, wrapper)
	}

	filesWrapper := &Wrappers{
		collection:          &files,
		directories:         nil,
		files:               nil,
		Error:               errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}

	filesWrapper.files = filesWrapper
	wrappers.files = filesWrapper

	return wrappers.files
}

func (wrappers *Wrappers) RootDirs() *Wrappers {
	if wrappers.directories != nil {
		return wrappers.directories
	}

	if wrappers.IsEmpty() {
		wrappers.directories = EmptyWrappers()
	}

	dirs := make(
		[]Wrapper,
		0,
		wrappers.Length())

	for _, wrapper := range *wrappers.collection {
		if !wrapper.IsDirectory {
			continue
		}

		dirs = append(dirs, wrapper)
	}

	dirWrappers := &Wrappers{
		collection:          &dirs,
		directories:         nil,
		files:               nil,
		Error:               errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}

	dirWrappers.directories = dirWrappers
	wrappers.directories = dirWrappers

	return wrappers.directories
}

func (wrappers *Wrappers) PathsCollection() *PathsCollection {
	if wrappers.pathsCollection != nil {
		return wrappers.pathsCollection
	}

	wrappers.pathsCollection = NewPathsUsingWrappers(
		wrappers.RootPath,
		wrappers.separator,
		wrappers)

	return wrappers.pathsCollection
}

func (wrappers *Wrappers) FileNamesCollection() *FileNamesCollection {
	if wrappers.fileNamesCollection != nil {
		return wrappers.fileNamesCollection
	}

	wrappers.fileNamesCollection = NewFileNamesUsingWrappers(wrappers)

	return wrappers.fileNamesCollection
}

func (wrappers *Wrappers) IsEmpty() bool {
	return wrappers.Error.HasError() ||
		wrappers.collection == nil ||
		len(*wrappers.collection) == 0
}

func (wrappers *Wrappers) Collection() *[]Wrapper {
	return wrappers.collection
}

func (wrappers *Wrappers) Length() int {
	if wrappers.collection == nil || *wrappers.collection == nil {
		return 0
	}

	return len(*wrappers.collection)
}

func (wrappers *Wrappers) IsPathContains(
	path string,
	isCaseSensitive bool,
) bool {
	return wrappers.
		PathsCollection().
		IsContains(
			path,
			isCaseSensitive)
}

func (wrappers *Wrappers) IsNameContains(
	name string,
	isCaseSensitive bool,
) bool {
	return wrappers.
		FileNamesCollection().
		IsContains(
			name,
			isCaseSensitive)
}
