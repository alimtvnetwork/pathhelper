package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Pictures directory path. If directory doesn't exist, then function creates the directory and returns the path as a string.
func GetUserPicturesPath() string {
	return GetUserPathOf(enums.Pictures.Value())
}
