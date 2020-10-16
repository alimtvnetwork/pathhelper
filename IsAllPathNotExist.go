package pathhelper

// returns false if any of the paths exists
func IsAllPathNotExist(paths ...string) bool {
	for _, path := range paths {
		if IsPathExist(path) {
			return false
		}
	}

	return true
}
