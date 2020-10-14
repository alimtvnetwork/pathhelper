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

func NormalizePathUsingSeparator(pathSeparator, givenPath string) string {
	if pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	firstStepNormalize := GetCompiledPath(givenPath, &normalizeMap)

	return RemoveAndFixDoubleSeparatorToFinalSeparator(pathSeparator, strings.TrimSpace(firstStepNormalize))
}
