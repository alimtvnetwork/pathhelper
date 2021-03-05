package ispaths

func AllDirectories(fullPaths ...string) bool {
	if fullPaths == nil {
		return false
	}

	return AllDirectoriesPtr(&fullPaths)
}
