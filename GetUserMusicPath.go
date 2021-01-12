package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Music directory path as a string.
func GetUserMusicPath() string {
	return enums.Music.CombineWith(GetUserPath())
}
