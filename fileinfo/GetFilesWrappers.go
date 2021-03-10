package fileinfo

import (
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func GetFilesWrappers(rootPath string, wrapperIn *[]Wrapper) *Wrappers {
	if wrapperIn == nil || *wrapperIn == nil {
		return EmptyWrappers()
	}

	length := len(*wrapperIn)
	files := make([]Wrapper, 0, length)

	for _, wrapper := range *wrapperIn {
		if !wrapper.IsFile {
			continue
		}

		files = append(files, wrapper)
	}

	filesWrappers := &Wrappers{
		RootPath:            rootPath,
		collection:          &files,
		directories:         EmptyWrappers(),
		files:               nil,
		recursiveDirs:       nil,
		Error:               errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}

	filesWrappers.files = filesWrappers

	return filesWrappers
}
