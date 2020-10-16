package pathhelper

// Returns true if any of the paths exists
func IsAnyPathExists(paths ...string) bool {
	for _, path := range paths {
		if IsPathExist(path) {
			return true
		}
	}

	return false
}
