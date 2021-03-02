package pathhelper

import "gitlab.com/evatix-go/pathhelper/ispath"

// returns false if any of the path not exists
func IsAllPathExist(paths ...string) bool {
	for _, path := range paths {
		if !ispath.Exist(path) {
			return false
		}
	}

	return true
}
