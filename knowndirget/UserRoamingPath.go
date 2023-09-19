package knowndirget

import (
	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns Roaming directory path as a string.
func UserRoamingPath() string {
	return knowndir.Roaming.CombineWith(UserPath())
}
