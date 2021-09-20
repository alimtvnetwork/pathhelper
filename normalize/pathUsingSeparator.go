package normalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/osconsts"
)

// pathUsingSeparator
//
// Always returns path without the separator at the end.
// Separator must be one char.
// Long path fix will not be applied other than Windows operating system.
func pathUsingSeparator(
	isLongPathFix bool,
	isForceLongPath bool,
	pathSeparator,
	givenPath string,
) string {
	if givenPath == "" {
		return givenPath
	}

	finalResult := removeAndFixDoubleSeparatorToFinalSeparator(
		pathSeparator,
		givenPath)
	if len(finalResult) == 0 {
		return finalResult
	}

	isApplyLongPathFix := isLongPathFix && osconsts.IsWindows
	if isApplyLongPathFix && finalResult[constants.Zero] == pathSeparator[constants.Zero] {
		finalResult = finalResult[constants.One:]
	}

	if isApplyLongPathFix {
		finalResult = getLongPathFixedUsingSeparator(
			isForceLongPath,
			pathSeparator,
			finalResult,
		)
	}

	lastIndex := len(finalResult) - 1

	if finalResult[lastIndex] == pathSeparator[coreindexes.First] {
		// removing last path separator
		return finalResult[:lastIndex]
	}

	return finalResult
}
