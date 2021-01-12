package pathhelper

import "path/filepath"

// Returns path to Users directory
func GetUsersPath() string {
	return filepath.Dir(GetUserPath())
}
