package pathhelper

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var normalizeMap = map[string]string{
	constants.UriSchemePrefixStandard:   "",
	constants.UriSchemePrefixTwoSlashes: "",
	constants.DoubleForwardSlash:        constants.PathSeparator,
	constants.DoubleBackSlash:           constants.PathSeparator,
}

func NormalizePath(givenPath string) string {
	if pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	// if givenPath does not contain file prefixes
	if !HasPathIssues(givenPath) {
		return path.Clean(givenPath)
	}

	normalize := GetCompiledPath(givenPath, &normalizeMap)

	return RemovingDoubleSeparator(strings.TrimSpace(normalize))
}
