package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
)

// By default apply long path fix and regular normalize using os.PathSeparator
func NormalizePath(givenPath string) string {
	return NormalizePathUsingSeparator(constants.PathSeparator, givenPath, true)
}
