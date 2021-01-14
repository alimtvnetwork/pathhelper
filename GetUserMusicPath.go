package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Music directory path as a string.
func GetUserMusicPath() string {
	return knowndir.Music.CombineWith(GetUserPath())
}
