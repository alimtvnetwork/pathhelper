package pathhelper

import "gitlab.com/evatix-go/pathhelper/pathhelpercore"

func NormalizePathUsingSeparatorIf(isNormalize bool, pathSeparator, givenPath string) string {
	if !isNormalize || pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	return NormalizePathUsingSeparator(pathSeparator, givenPath)
}
