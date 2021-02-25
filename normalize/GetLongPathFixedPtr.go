package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

// Adds constants.LongPathQuestionMarkPrefix if path is longer than 255 and doesn't already contains it.
// Ignores prefix apply if already has it (constants.LongPathUncPrefix or constants.LongPathQuestionMarkPrefix ) or path is empty or length less than 255
// if path starts with `\\` then replaces with constants.LongPathUncPrefix
func GetLongPathFixedPtr(givenAbsolutePath *string, isForce bool) string {
	isIgnoreCase :=
		IsEmptyPathPtr(givenAbsolutePath) ||
			len(*givenAbsolutePath) < 255

	if (isIgnoreCase && !isForce) || osconsts.IsUnixGroup {
		return *givenAbsolutePath
	}

	currentPath := *givenAbsolutePath

	hasLongPathPrefixAlready := strings.HasPrefix(currentPath, constants.LongPathUncPrefix) ||
		strings.HasPrefix(currentPath, constants.LongPathQuestionMarkPrefix)

	if hasLongPathPrefixAlready {
		return currentPath
	}

	// Broken fix
	if strings.HasPrefix(currentPath, consts.BrokenLongPathQuestionMarkPrefix) {
		return strings.Replace(
			currentPath,
			consts.BrokenLongPathQuestionMarkPrefix,
			constants.LongPathQuestionMarkPrefix,
			constants.One)
	}

	// Broken fix
	if strings.HasPrefix(currentPath, consts.BrokenLongPathUncPrefix) {
		return strings.Replace(
			currentPath,
			consts.BrokenLongPathUncPrefix,
			constants.LongPathUncPrefix,
			constants.One)
	}

	if strings.HasPrefix(currentPath, constants.DoubleBackSlash) {
		return strings.Replace(
			currentPath,
			constants.DoubleBackSlash,
			constants.LongPathUncPrefix,
			constants.One)
	}

	return constants.LongPathQuestionMarkPrefix + currentPath
}
