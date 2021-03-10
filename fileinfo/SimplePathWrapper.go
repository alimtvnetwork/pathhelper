package fileinfo

type SimplePathWrapper struct {
	Path        string
	IsDirectory bool
}

func (pathWrapper *SimplePathWrapper) IsEquals(anotherWrapper SimplePathWrapper) bool {
	return pathWrapper.IsDirectory == anotherWrapper.IsDirectory &&
		pathWrapper.Path == anotherWrapper.Path
}

func (pathWrapper *SimplePathWrapper) IsEqualsPtr(anotherWrapper *SimplePathWrapper) bool {
	if anotherWrapper == nil {
		return false
	}

	if pathWrapper == anotherWrapper {
		return true
	}

	return pathWrapper.IsDirectory == anotherWrapper.IsDirectory &&
		pathWrapper.Path == anotherWrapper.Path
}

func (pathWrapper *SimplePathWrapper) String() string {
	return pathWrapper.Path
}
