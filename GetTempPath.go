package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns temp directory. After checking in on all possible locations, if directory doesn't exist, then doesnt create the directory
// and returns the path as a string.
func GetTempPath() string {
	if IsWindows() {
		return os.Getenv(enums.Temp.Value())
	}

	desiredTempPathUnix := enums.TempDir.CombineWith(GetUserPath())

	// Checking if temp directory is available
	tempUnix := os.Getenv(enums.TempDir.Value())

	if IsAllPathNotExist(desiredTempPathUnix, tempUnix) {
		return desiredTempPathUnix
	}

	return tempUnix
}
