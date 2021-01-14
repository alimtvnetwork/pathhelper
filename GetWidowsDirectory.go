package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns windows directory path.
func GetWidowsDirectory() string {
	return os.Getenv(knowndir.WindowsDirectory.Value())
}
