package pathhelper

import "gitlab.com/evatix-go/pathhelper/knowndir"

// Returns Windows root as a string
func GetWindowsRoot() string {
	return knowndir.WindowsCDrive.Value()
}
