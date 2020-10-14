package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
)

func NormalizePath(givenPath string) string {
	return NormalizePathUsingSeparator(constants.PathSeparator, givenPath)
}
