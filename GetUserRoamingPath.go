package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Roaming directory path as a string.
func GetUserRoamingPath() string {
	return enums.Roaming.GetPrefixCombinedWith(GetUserPath())
}
