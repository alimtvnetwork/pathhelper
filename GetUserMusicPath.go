package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Music directory path. If directory doesn't exist, then function creates the directory and returns the path as a string.
func GetUserMusicPath() string {
	return GetUserPathOf(enums.Music.Value())
}
