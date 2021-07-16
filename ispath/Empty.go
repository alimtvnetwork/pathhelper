package ispath

func Empty(path string) bool {
	return len(path) == 0
}

func EmptyPtr(path *string) bool {
	return path == nil || Empty(*path)
}
