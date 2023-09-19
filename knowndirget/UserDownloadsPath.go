package knowndirget

import (
	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns downloads directory path as a string.
func UserDownloadsPath() string {
	return knowndir.Downloads.CombineWith(UserPath())
}
