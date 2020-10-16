package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Videos directory path. If directory doesn't exist, then function creates the directory and returns the path as a string.
func GetUserVideosPath() string {
	return GetUserPathOf(enums.Videos.Value())
}
