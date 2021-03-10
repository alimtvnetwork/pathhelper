package fileinfo

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

type Wrappers struct {
	RootPath            string
	Separator           string
	Items               *[]*Wrapper
	directories         *Wrappers
	files               *Wrappers
	ErrorWrapper        *errorwrapper.Wrapper
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
			wrappers.Separator,
			wrappers.RootPath,
			isErrorContinueDefault)

	errorCollection.HandleError()

	wrappers.recursiveDirs = NewPathsUsingPaths(
		wrappers.RootPath,
		wrappers.Separator,
		allPaths)

	wrappers.recursiveDirs.parentWrappers = wrappers
	wrappers.recursiveDirs.directories = allPaths

	return wrappers.recursiveDirs
}

func (wrappers *Wrappers) RecursivePathsFilter(
	filter pathfuncs.Filter,
) *PathsCollection {
	allPaths, errorCollection := recursiveinternal.GetFilterPaths(
		wrappers.Separator,
		wrappers.RootPath,
		isErrorContinueDefault,
		filter)

	errorCollection.HandleError()

	newPathsCollection := NewPathsUsingPaths(
		wrappers.RootPath,
		wrappers.Separator,
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
		[]*Wrapper,
		0,
		wrappers.Length())

	for _, wrapper := range *wrappers.Items {
		if !wrapper.IsFile {
			continue
		}

		files = append(files, wrapper)
	}

	filesWrapper := &Wrappers{
		Items:               &files,
		directories:         nil,
		files:               nil,
		ErrorWrapper:        errnew.EmptyPtr,
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
		[]*Wrapper,
		0,
		wrappers.Length())

	for _, wrapper := range *wrappers.Items {
		if !wrapper.IsDirectory {
			continue
		}

		dirs = append(dirs, wrapper)
	}

	dirWrappers := &Wrappers{
		Items:               &dirs,
		directories:         nil,
		files:               nil,
		ErrorWrapper:        errnew.EmptyPtr,
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
		wrappers.Separator,
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
	return wrappers.ErrorWrapper.HasError() ||
		wrappers.Items == nil ||
		len(*wrappers.Items) == 0
}

func (wrappers *Wrappers) Length() int {
	if wrappers.Items == nil || *wrappers.Items == nil {
		return 0
	}

	return len(*wrappers.Items)
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
