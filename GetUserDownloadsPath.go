package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns downloads directory path. If directory doesn't exist, then function creates the directory and returns the path as a string.
func GetUserDownloadsPath() string {
	return GetUserPathOf(enums.Downloads.Value())
}
