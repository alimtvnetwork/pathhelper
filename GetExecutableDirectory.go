package pathhelper

import (
	"path/filepath"
)

func GetExecutableDirectory() string {
	exePath := GetExecutablePath()
	exeDir := filepath.Dir(exePath)

	return exeDir
}
