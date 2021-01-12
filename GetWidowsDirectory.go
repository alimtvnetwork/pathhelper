package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns windows directory path.
func GetWidowsDirectory() string {
	return os.Getenv(enums.WindowsDirectory.Value())
}
