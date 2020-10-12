package pathhelper

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var normalizeMap = map[string]string{
	constants.UriSchemePrefixStandard: "",
}

func NormalizePath(givenPath string) string {
	if pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	// if givenPath does not contain file prefixes
	if !HasPathIssues(givenPath) {
		return path.Clean(givenPath)
	}

	// when givenPath contains UriSchemePrefixStandard
	pathWithoutPrefix := strings.Replace(givenPath, constants.UriSchemePrefixStandard, "", 1)

	// when givenPath contains UriSchemePrefixTwoSlashes
	pathWithoutPrefix = strings.Replace(pathWithoutPrefix, constants.UriSchemePrefixTwoSlashes, "", 1)

	// removing doubles
	pathWithoutDouble := RemovingDouble(pathWithoutPrefix)

	// replacing with correct separator
	var pathWithCorrectSeparator string

	if IsWindows() {
		pathWithCorrectSeparator = strings.ReplaceAll(pathWithoutDouble, constants.ForwardSlash, constants.BackSlash)
	} else {
		pathWithCorrectSeparator = strings.ReplaceAll(pathWithoutDouble, constants.BackSlash, constants.ForwardSlash)
	}

	return RemovingDouble(strings.TrimSpace(pathWithCorrectSeparator))
}
