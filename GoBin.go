package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/constants"
)

// Returns env go bin path
func GoBin() string {
	return os.Getenv(constants.GoBinPath)
}
