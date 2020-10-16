package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to System32 path in windows system.
func GetSystem32() string {
	return os.Getenv(enums.System32.Value())
}
