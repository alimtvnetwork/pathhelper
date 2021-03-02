package pathhelper

import "gitlab.com/evatix-go/pathhelper/ispath"

// Returns true if any of the paths exists
func IsAnyPathExists(paths ...string) bool {
	for _, path := range paths {
		if ispath.Exist(path) {
			return true
		}
	}

	return false
}
