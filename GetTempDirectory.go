package pathhelper

import "os"

func GetTempDirectory() string {
	return os.TempDir()
}
