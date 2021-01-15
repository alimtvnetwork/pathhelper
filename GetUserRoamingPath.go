package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Roaming directory path as a string.
func GetUserRoamingPath() string {
	return knowndir.Roaming.CombineWith(GetUserPath())
}
