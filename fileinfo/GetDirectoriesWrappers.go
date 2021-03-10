package fileinfo

import (
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func GetDirectoriesWrappers(rootPath string, wrapperIn *[]Wrapper) *Wrappers {
	if wrapperIn == nil || *wrapperIn == nil {
		return EmptyWrappers()
	}

	length := len(*wrapperIn)
	dirs := make([]Wrapper, 0, length)

	for _, wrapper := range *wrapperIn {
		if !wrapper.IsDirectory {
			continue
		}

		dirs = append(dirs, wrapper)
	}

	dirWrappers := &Wrappers{
		RootPath:            rootPath,
		collection:          &dirs,
		directories:         nil,
		files:               EmptyWrappers(),
		recursiveDirs:       nil,
		Error:               errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}

	dirWrappers.directories = dirWrappers

	return dirWrappers
}
