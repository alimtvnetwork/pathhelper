package pathhelper

import (
	"path/filepath"
)

// Represents the directory where the application is running from.
func GetExecutableDirectory() string {
	exePath := GetExecutablePath()
	exeDir := filepath.Dir(exePath)

	return exeDir
}
