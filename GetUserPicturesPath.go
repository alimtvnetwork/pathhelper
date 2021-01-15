package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Pictures directory path as a string.
func GetUserPicturesPath() string {
	return knowndir.Pictures.CombineWith(GetUserPath())
}
