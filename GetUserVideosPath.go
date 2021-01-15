package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Videos directory path as a string.
func GetUserVideosPath() string {
	return knowndir.Videos.CombineWith(GetUserPath())
}
