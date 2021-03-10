package fileinfo

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

type SimplePathWrapper struct {
	Path        string
	IsDirectory bool
}

func (simplePathWrapper *SimplePathWrapper) IsEquals(anotherWrapper SimplePathWrapper) bool {
	return simplePathWrapper.IsDirectory == anotherWrapper.IsDirectory &&
		simplePathWrapper.Path == anotherWrapper.Path
}

func (simplePathWrapper *SimplePathWrapper) IsEqualsPtr(anotherWrapper *SimplePathWrapper) bool {
	if anotherWrapper == nil {
		return false
	}

	if simplePathWrapper == anotherWrapper {
		return true
	}

	return simplePathWrapper.IsDirectory == anotherWrapper.IsDirectory &&
		simplePathWrapper.Path == anotherWrapper.Path
}

func (simplePathWrapper *SimplePathWrapper) String() string {
	return simplePathWrapper.Path
}

func (simplePathWrapper *SimplePathWrapper) BaseDir() string {
	return splitinternal.GetBaseDir(simplePathWrapper.Path)
}

func (simplePathWrapper *SimplePathWrapper) GetFileNamePlusExt() (filename, ext string) {
	return splitinternal.GetFilenamePlusExt(simplePathWrapper.Path)
}

func (simplePathWrapper *SimplePathWrapper) GetBothExtension() (dotExt, ext string) {
	return splitinternal.GetBothExtension(simplePathWrapper.Path)
}

func (simplePathWrapper *SimplePathWrapper) GetFileNameOnly() (fileNameWithoutExt string) {
	return splitinternal.GetFileNameWithoutExt(simplePathWrapper.Path)
}
