package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

// Always returns path without the separator at the end.
// Separator must be one char.
// Long path fix will not be applied other than Windows operating system.
func PathUsingSeparator(
	pathSeparator,
	givenPath string,
	isLongPathFix bool,
) string {
	if pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	firstStepNormalize := GetCompiledPath(givenPath, &normalizeMap)
	result := removeAndFixDoubleSeparatorToFinalSeparator(
		pathSeparator,
		strings.TrimSpace(firstStepNormalize))

	if isLongPathFix && osconsts.IsWindows {
		result = GetLongPathFixed(result)
	}

	if result[len(result)-1] == pathSeparator[0] {
		// removing last path separator
		return result[:len(result)-1]
	}

	return result
}
