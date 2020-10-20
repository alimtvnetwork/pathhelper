package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var normalizeMap = map[string]string{
	constants.UriSchemePrefixStandard:   "",
	constants.UriSchemePrefixTwoSlashes: "",
}

func NormalizePathUsingSeparator(pathSeparator, givenPath string, isLongPathFix bool) string {
	if pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	firstStepNormalize := GetCompiledPath(givenPath, &normalizeMap)

	result := RemoveAndFixDoubleSeparatorToFinalSeparator(pathSeparator, strings.TrimSpace(firstStepNormalize))

	if isLongPathFix {
		result = GetLongPathFixed(result)
	}

	return result
}
