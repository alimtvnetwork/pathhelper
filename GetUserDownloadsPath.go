package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns downloads directory path as a string.
func GetUserDownloadsPath() string {
	return knowndir.Downloads.CombineWith(GetUserPath())
}
