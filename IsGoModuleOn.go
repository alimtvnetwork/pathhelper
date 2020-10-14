package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/constants"
)

// Returns env go bin path
func IsGoModuleOn() bool {
	return os.Getenv(constants.Go111ModuleEnvironment) == constants.On
}
