package normalize

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/osconsts"
)

func LongPathFixPlusClean(
	isForceLongPathFix bool,
	pathSeparator,
	givenPath string,
) string {
	if givenPath == "" {
		return givenPath
	}

	finalResult := filepath.Clean(givenPath)
	if len(finalResult) == 0 {
		return finalResult
	}

	if osconsts.IsWindows && finalResult[constants.Zero] == pathSeparator[constants.Zero] {
		finalResult = finalResult[constants.One:]
	}

	if osconsts.IsWindows {
		finalResult = getLongPathFixedUsingSeparator(
			pathSeparator,
			finalResult,
			isForceLongPathFix)
	}

	length := len(finalResult)
	lastIndex := length - 1

	if finalResult[lastIndex] == pathSeparator[coreindexes.First] {
		// removing last path separator
		return finalResult[:lastIndex]
	}

	return finalResult
}
