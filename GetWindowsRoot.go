package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns Windows root as a string
func GetWindowsRoot() string {
	return enums.WindowsCDrive.Value()
}
