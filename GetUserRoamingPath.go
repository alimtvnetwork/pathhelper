package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Roaming directory path. If directory doesn't exist, then function creates the directory and returns the path as a string.
func GetUserRoamingPath() string {
	return GetUserPathOf(enums.Roaming.Value())
}
