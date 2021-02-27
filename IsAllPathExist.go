package pathhelper

// returns false if any of the path not exists
func IsAllPathExist(paths ...string) bool {
	for _, path := range paths {
		if !IsPathExist(path) {
			return false
		}
	}

	return true
}
