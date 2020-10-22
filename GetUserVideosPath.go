package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns Videos directory path as a string.
func GetUserVideosPath() string {
	return enums.Videos.GetPrefixCombinedWith(GetUserPath())
}
