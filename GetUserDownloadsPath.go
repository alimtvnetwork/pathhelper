package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns downloads directory path as a string.
func GetUserDownloadsPath() string {
	return enums.Downloads.GetPrefixCombinedWith(GetUserPath())
}
