package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Pictures directory path as a string.
func GetUserPicturesPath() string {
	return enums.Pictures.GetPrefixCombinedWith(GetUserPath())
}
