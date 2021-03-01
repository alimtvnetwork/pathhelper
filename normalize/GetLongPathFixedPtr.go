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
func GetLongPathFixedPtr(
	separator string,
	givenAbsolutePath string,
	isForce bool,
) string {
	isIgnoreCase :=
		IsEmptyPath(givenAbsolutePath) ||
			len(givenAbsolutePath) < 255

	isWindowsSeparator := separator ==
		constants.WindowsPathSeparator

	if !isWindowsSeparator {
		return givenAbsolutePath
	}

	isForce = isForce && isWindowsSeparator

	if (isIgnoreCase && !isForce) || osconsts.IsUnixGroup {
		return givenAbsolutePath
	}

	hasLongPathPrefixAlready :=
		strings.HasPrefix(givenAbsolutePath, constants.LongPathUncPrefix) ||
			strings.HasPrefix(givenAbsolutePath, constants.LongPathQuestionMarkPrefix)

	if hasLongPathPrefixAlready {
		return givenAbsolutePath
	}

	// Broken fix
	if strings.HasPrefix(givenAbsolutePath, consts.BrokenLongPathQuestionMarkPrefix) {
		return strings.Replace(
			givenAbsolutePath,
			consts.BrokenLongPathQuestionMarkPrefix,
			constants.LongPathQuestionMarkPrefix,
			constants.One)
	}

	// Broken fix
	if strings.HasPrefix(givenAbsolutePath, consts.BrokenLongPathUncPrefix) {
		return strings.Replace(
			givenAbsolutePath,
			consts.BrokenLongPathUncPrefix,
			constants.LongPathUncPrefix,
			constants.One)
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
