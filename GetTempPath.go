package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
	"os"
)

// Returns temp directory. After checking in on all possible locations, if directory doesn't exist, then creates the directory
// and returns the path as a string.
func GetTempPath() string {
	if IsWindows() {
		return os.Getenv(enums.Temp.Value())
	}

	desiredTempPathUnix := GetCombinePathWith(GetUserPath(), enums.TempDir.Value())

	// Checking if temp directory is available
	tempUnix := os.Getenv(enums.TempDir.Value())

	if IsAllPathNotExist(desiredTempPathUnix, tempUnix) {
		CreateDirectory(desiredTempPathUnix, constants.Perm)

		return desiredTempPathUnix
	}

	return tempUnix
}
