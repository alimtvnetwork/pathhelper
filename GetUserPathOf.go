package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
)

// Takes in input of a directory under users directory and returns its absolute path.
func GetUserPathOf(directoryName string) string {
	outputPath := GetCombinePathWith(GetUserPath(), directoryName)

	CreateDirectory(outputPath, constants.Perm)

	return outputPath
}
