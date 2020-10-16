package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to %AppData% in windows. If directory doesn't exist it creates the directory and returns the path as a string.
// Requires further investigation for linux platform
func GetAppDataPath() string {
	appDataPath := os.Getenv(enums.AppData.Value())

	CreateDirectory(appDataPath, constants.Perm)

	return appDataPath
}
