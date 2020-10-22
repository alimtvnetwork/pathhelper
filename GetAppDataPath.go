package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to %AppData% in windows. If directory doesn't exist it still returns the path as a string.
// Requires further investigation for linux platform: https://stackoverflow.com/questions/1510104/where-to-store-application-data-non-user-specific-on-linux,
// https://stackoverflow.com/questions/17517131/appdata-in-non-windows
func GetAppDataPath() string {
	if IsWindows() {
		return os.Getenv(enums.AppData.Value())
	}

	return enums.AppDataUnix.Value()
}
