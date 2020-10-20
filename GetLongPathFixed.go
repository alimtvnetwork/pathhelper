package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix ) or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
func GetLongPathFixed(givenAbsolutePath string) string {
	isIgnoreCase := pathhelpercore.IsEmptyPath(givenAbsolutePath) || len(givenAbsolutePath) < 255

	if isIgnoreCase {
		return givenAbsolutePath
	}

	hasLongPathPrefixAlready := strings.HasPrefix(givenAbsolutePath, constants.LongPathUncPrefix) ||
		strings.HasPrefix(givenAbsolutePath, constants.LongPathQuestionMarkPrefix)

	if hasLongPathPrefixAlready {
		return givenAbsolutePath
	}

	if strings.HasPrefix(givenAbsolutePath, constants.DoubleBackSlash) {
		return strings.Replace(
			givenAbsolutePath,
			constants.DoubleBackSlash,
			constants.LongPathUncPrefix,
			constants.One)
	}

	return constants.LongPathQuestionMarkPrefix + givenAbsolutePath
}
