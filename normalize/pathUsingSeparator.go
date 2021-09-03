package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/osconsts"
)

// pathUsingSeparator Always returns path without the separator at the end.
// Separator must be one char.
// Long path fix will not be applied other than Windows operating system.
func pathUsingSeparator(
	isLongPathFix bool,
	isForceLongPath bool,
	pathSeparator,
	givenPath string,
) string {
	if givenPath == "" || osconsts.IsUnixGroup {
		return givenPath
	}

	firstStepNormalize := GetCompiledPath(
		givenPath,
		normalizeMap)
	result := removeAndFixDoubleSeparatorToFinalSeparator(
		pathSeparator,
		strings.TrimSpace(firstStepNormalize))

	if isLongPathFix && osconsts.IsWindows {
		result = getLongPathFixedUsingSeparator(
			pathSeparator,
			result,
			isForceLongPath)
	}

	length := len(result)
	lastIndex := length - 1

	if result[lastIndex] == pathSeparator[coreindexes.First] {
		// removing last path separator
		return result[:lastIndex]
	}

	return result
}
