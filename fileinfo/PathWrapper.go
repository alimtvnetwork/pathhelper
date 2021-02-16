package fileinfo

type PathWrapper struct {
	Path        string
	IsDirectory bool
}

func (pathWrapper *PathWrapper) IsEquals(anotherWrapper PathWrapper) bool {
	return pathWrapper.IsDirectory == anotherWrapper.IsDirectory &&
		pathWrapper.Path == anotherWrapper.Path
}

func (pathWrapper *PathWrapper) IsEqualsPtr(anotherWrapper *PathWrapper) bool {
	if anotherWrapper == nil {
		return false
	}

	if pathWrapper == anotherWrapper {
		return true
	}

	return pathWrapper.IsDirectory == anotherWrapper.IsDirectory &&
		pathWrapper.Path == anotherWrapper.Path
}

func (pathWrapper *PathWrapper) String() string {
	return pathWrapper.Path
}
